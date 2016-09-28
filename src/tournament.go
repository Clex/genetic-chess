package geneticchess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type Game struct {
	players []*AI
}

type Tournament struct {
	players       []*AI
	games         []*Game
	gamesByPlayer int
}

type Result struct {
	score  float64
	Player *AI
}
type Results []Result

func (r Results) Len() int {
	return len(r)
}

func (r Results) Less(i int, j int) bool {
	return r[i].score >= r[j].score
}

func (r Results) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

func NewTournament(players []*AI, nbElites uint, nbChild uint, nbGames uint,
	nbMutations uint, mutationSize float64) *Tournament {
	if players == nil || len(players) == 0 {
		panic("need at least one ai")
	}

	for i := uint(1); uint(len(players)) < nbElites; i++ {
		clone := players[0].mute(uint(len(genes)), 1.0)
		players = append(players, clone)
	}

	t := &Tournament{}

	for i := 0; i < int(nbChild); i++ {
		for _, parent := range players {
			child := parent.mute(nbGames, mutationSize)
			t.players = append(t.players, child)
		}
	}

	t.players = append(t.players, players...)

	n := len(t.players)
	t.gamesByPlayer = (n - 1) * int(nbGames) * 2

	for _, player1 := range t.players {
		for _, player2 := range t.players {
			if player1 == player2 {
				continue
			}

			for i := uint(0); i < nbGames/2; i++ {
				t.games = append(t.games, &Game{
					players: []*AI{player1, player2},
				}, &Game{
					players: []*AI{player2, player1},
				})
			}
		}
	}

	return t
}

func (t *Tournament) Play(timeToThink time.Duration,
	maxDepth uint, nbParallelGames uint, verbose bool) Results {
	start := time.Now()

	n := 0
	playingGames := uint(0)
	playedGames := 0

	resChan := make(chan [2]*Result)
	scores := make(map[*AI]float64)

	if verbose == true {
		fmt.Println("Starting new tournament")
	}

	for {
		if verbose == true {
			fmt.Printf("\rPlaying game %d of %d...", playedGames+1, len(t.games))
		}

		if n < len(t.games) && playingGames < nbParallelGames {
			game := t.games[n]
			go game.players[0].Play(game.players[1], White,
				timeToThink, maxDepth, resChan)
			playingGames++
			n++
			continue
		}

		res := <-resChan
		scores[res[0].Player] += res[0].score
		scores[res[1].Player] += res[1].score
		playingGames--
		playedGames++
		if playedGames == len(t.games) {
			break
		}
	}

	if verbose == true {
		diff := time.Now().Sub(start)
		fmt.Println("")
		fmt.Printf("\nTournament finished in %v (%v/game)\n",
			diff, diff/time.Duration(len(t.games)))
	}

	var res Results

	for AI, score := range scores {
		res = append(res, Result{
			Player: AI,
			score:  score,
		})
	}

	sort.Sort(res)

	if verbose == true {
		fmt.Println("Results:")
		for _, r := range res {
			fmt.Printf("%s: %.2f/%d\n", r.Player.String(),
				r.score, t.gamesByPlayer)
		}
		fmt.Println("")
	}

	return res
}

func (t *Tournament) displayFileUpdated(quiet bool,
	exChamp *AI, champ *AI, file string) bool {
	if quiet == true {
		return false
	}

	if exChamp == champ {
		return false
	}

	fmt.Printf("\nEx-champion %s is replaced by %s in %s\n",
		exChamp.String(), champ.String(), file)

	return true
}

func RunTournaments(file string, timeToThink time.Duration, maxDepth uint,
	nbQualified uint, nbChildren uint, nbGames uint, nbMutations uint,
	mutationSize float64, nbParallelGames uint, rounds uint, quiet bool) error {
	ai, err := NewAIFromFile(file)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		ai = NewAIRandom()
	}

	qualified := []*AI{ai}

	for i := uint(0); ; i++ {
		if rounds > 0 && i == rounds {
			break
		}

		exChamp := qualified[0]
		t := NewTournament(qualified, nbQualified,
			nbChildren, nbGames, nbMutations, mutationSize)
		res := t.Play(timeToThink, maxDepth, nbParallelGames, !quiet)

		if rounds > 0 && i+1 == rounds {
			fmt.Printf("Winner: %s", res[0].Player.String())
			qualified = []*AI{res[0].Player}
			updated := t.displayFileUpdated(
				quiet, exChamp, res[0].Player, file)
			if updated == false {
				fmt.Println("")
			}
		} else {
			fmt.Printf("Qualified for next round: ")
			qualified = []*AI{}
			for i := uint(0); i < nbQualified; i++ {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s", res[i].Player.String())
				qualified = append(qualified, res[i].Player)
			}

			updated := t.displayFileUpdated(
				quiet, exChamp, res[0].Player, file)
			if updated == true {
				fmt.Println("")
			} else {
				fmt.Println("\n")
			}
		}

		// Save champion.
		winnerJSON, err := json.Marshal(qualified[0])
		if err != nil {
			return fmt.Errorf("cannot marshal json: %v", err)
		}

		err = ioutil.WriteFile(file, winnerJSON, 0644)
		if err != nil {
			return fmt.Errorf("cannot write %s: %v", file, err)
		}
	}

	return nil
}
