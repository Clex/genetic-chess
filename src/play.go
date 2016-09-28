package geneticchess

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Play(file string, timeToThink time.Duration, maxDepth uint) error {
	ai, err := NewAIFromFile(file)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	board := NewBoard()

	for i := 0; ; i++ {
		if i > 0 {
			fmt.Println("")
		}

		board.Dump()

		fmt.Print("Move: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("cannot read string: %v", err)
		}

		text = strings.TrimSpace(text)
		parts := strings.Split(text, ":")
		if len(parts) != 2 {
			fmt.Printf("invalid move: %s\n", text)
			continue
		}

		from, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("invalid origin square: %s\n", parts[0])
			continue
		}

		dest := parts[1]

		if len(dest) == 0 {
			fmt.Println("invalid destination square")
			continue
		}

		var promotion string

		to, err := strconv.Atoi(dest)
		if err != nil {
			promotion = dest[len(dest)-1:]
			dest = dest[:len(dest)-1]
			to, err = strconv.Atoi(dest)
			if err != nil {
				fmt.Printf("invalid destination square: %s\n", dest)
				continue
			}
		}

		move := &Move{
			from:      Position(from),
			to:        Position(to),
			promoteTo: StringToPieceType(promotion),
		}

		m := board.GetMoves()
		ok := false
		for _, possibleMove := range m {
			if possibleMove.Equals(move) {
				ok = true
			}
		}
		if !ok {
			fmt.Printf("move %s not allowed\n", move.String())
			continue
		}

		state := board.Move(move)
		if state != StatePlaying {
			fmt.Println(state)
			break
		}

		board.Dump()
		fmt.Println("looking for best move")
		move = ai.GetBestMove(board, timeToThink, maxDepth)
		fmt.Printf("best move found: %s\n\n", move.String())
		state = board.Move(move)
		if state != StatePlaying {
			fmt.Println(state)
			break
		}
	}

	return nil
}
