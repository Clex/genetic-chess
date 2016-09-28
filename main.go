package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	gc "github.com/clex/genetic-chess/src"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	l := log.New(os.Stderr, "", 0)

	play := flag.Bool("play", false,
		"play against ai")
	file := flag.String("file", gc.DefaultFilePath,
		"data file, created if necessary")
	qualified := flag.Uint("qualified", 1,
		"number of ai qualified for the next tournament")
	games := flag.Uint("games", 2,
		"number of games to play against each other "+
			"(must be a multiple of 2 to keep white/black even)")
	children := flag.Uint("children", 2,
		"number of children for each qualified")
	mutations := flag.Uint("mutations", 1,
		"number of mutations between two generations")
	mutationSize := flag.Float64("mutation-size", 0.25,
		"maximum mutation size for a gene between two generations")
	timeToThink := flag.Duration("time-to-think", time.Millisecond*10,
		"maximum time to think for a move "+
			`(suffix with "ms", "s", "m" or "h"`)
	maxDepth := flag.Uint("max-depth", 3,
		"maximum allowed depth in games tree")
	parallelGames := flag.Uint("parallel-games", 1,
		"number of parallel games (each game uses a go routine)")
	rounds := flag.Uint("rounds", 0,
		"number of rounds (0 means infinite)")
	quiet := flag.Bool("quiet", false, "disables all output")

	flag.Parse()

	if *qualified == 0 {
		l.Fatalf("expected -qualified to be " +
			"a positive integer instead of 0")
	}
	if *games == 0 || *games%2 == 1 {
		l.Fatalf("expected -games to be " +
			"a positive integer instead of 0 and a multiple of 2")
	}
	if *children == 0 {
		l.Fatalf("expected -children to be " +
			"a positive integer instead of 0")
	}
	if *mutations == 0 {
		l.Fatalf("expected -mutations to be " +
			"a positive integer instead of 0")
	}
	if *mutationSize <= 0.01 || *mutationSize > 1.0 {
		l.Fatalf("expected -mutation-size to be "+
			"between 0.01 and 1.0 instead of %.2f",
			*mutationSize)
	}
	if *timeToThink <= 0 {
		l.Fatalf("expected -children to be "+
			"a positive duration instead of %v", *timeToThink)
	}
	if *maxDepth == 0 {
		l.Fatalf("expected -max-depth to be " +
			"a positive integer instead of 0")
	}
	if *parallelGames == 0 {
		l.Fatalf("expected -parallel-games to be " +
			"a positive integer instead of 0")
	}

	if *play == true {
		err := gc.Play(*file, *timeToThink, *maxDepth)
		if err != nil {
			l.Fatalf("cannot play: %v", err)
		}
	} else {
		err := gc.RunTournaments(*file, *timeToThink, *maxDepth,
			*qualified, *children, *games, *mutations, *mutationSize,
			*parallelGames, *rounds, *quiet)
		if err != nil {
			l.Fatalf("tournament failed: %v", err)
		}
	}
}
