package geneticchess

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAIEvalPieceValues(t *testing.T) {
	var info boardInfo

	ai := NewAI()
	b := NewBoard()

	score := ai.evalPieces(b, &info)
	if score != 0.0 {
		t.Fatalf("expected score %f instead of %f %s",
			0.0, score, fmt.Sprintf(""))
	}

	_ = b.Move(&Move{from: 52, to: 36})
	_ = b.Move(&Move{from: 11, to: 27})
	_ = b.Move(&Move{from: 36, to: 27})

	score = ai.evalPieces(b, &info)
	if score != 1.0 {
		t.Fatalf("expected score %f instead of %f", 1.0, score)
	}
}

func TestAIPieceValueBestMove(t *testing.T) {
	b := NewBoardFromDiagram(
		"    | | | | | | | | |" +
			"| | | | | | | | |" +
			"| | |N| |N| | | |" +
			"| | | |b| | | | |" +
			"| | |R| |P| | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"|k| | | | |K| | |")
	ai := NewAI()
	move := ai.GetBestMove(b, 0, 1)
	if move.from != 27 || move.to != 34 {
		t.Fatalf("expected move to be 27:34 instead of %d:%d",
			move.from, move.to)
	}
}

func TestAICheckmate(t *testing.T) {
	ai := NewAI()
	b := NewBoard()
	_ = b.Move(&Move{from: 52, to: 36})
	_ = b.Move(&Move{from: 12, to: 28})
	_ = b.Move(&Move{from: 59, to: 45})
	_ = b.Move(&Move{from: 8, to: 16})
	_ = b.Move(&Move{from: 61, to: 34})
	_ = b.Move(&Move{from: 16, to: 24})

	move := ai.GetBestMove(b, 0, 1)
	if move.from != 45 || move.to != 13 {
		t.Fatalf("expected move to be 45:13 instead of %d:%d",
			move.from, move.to)
	}
}

func TestAIBugKing(t *testing.T) {
	// A bug used to make finding a move from this position panic
	ai := NewAIEmpty()
	ai.Genes["PieceValuePawn"] = 10.0
	ai.Genes["PieceValueRook"] = 0.1
	ai.Genes["PieceValueKnight"] = 0.1
	ai.Genes["PieceValueBishop"] = 0.1
	ai.Genes["PieceValueQueen"] = 0.1
	b := NewBoardFromDiagram(
		"    | | | | | |B|R| |" +
			"| | | |r|P| |P| |" +
			"| | | | | |P| |P|" +
			"| | |q| |K| | | |" +
			"| | |b| | | | | |" +
			"| | | | | | | | |" +
			"| |p|p|p| |p|p|p|" +
			"| |n|b| | |k|n|r|")
	b.turn = Black

	move := ai.GetBestMove(b, 0, 1)
	_ = b.Move(move)
}

func TestAIBuildTree(t *testing.T) {
	ai := NewAI()
	b := NewBoardFromDiagram(
		"    | | | | | | | | |" +
			"| | | | | | | | |" +
			"|K| | | | | | | |" +
			"|p| | | | | | | |" +
			"|k| | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |")

	tree := ai.buildTree(b, 0, 1, false)
	if tree == nil {
		panic("")
	}
}

func TestAIGetBestMoveForcedCheckmate(t *testing.T) {
	ai := NewAI()
	ai.Genes["PruneRatio"] = 0.5

	tests := []struct {
		diagram string
		move    *Move
		turn    Color
	}{
		{"   | | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| |k| | | | | | |" +
			"| | | |n| | | | |" +
			"|K| | | |n| | | |",
			&Move{from: 60, to: 50}, White},
		{"   | | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| |r| | | | | | |" +
			"| | |k| | | | | |" +
			"| | | | | | | | |" +
			"|K| | | | | | | |",
			&Move{from: 42, to: 50}, White},
		{"   |k| | | |r| | | |" +
			"| | | | |n|r| | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | |K| |",
			&Move{from: 4, to: 6}, White},
		{"   | | | | | | | | |" +
			"| | | | | |p| |K|" +
			"| | | | | | |P|P|" +
			"| | | | | |P| | |" +
			"| | | | | | | | |" +
			"| |b|b| | | | | |" +
			"| | | | | | | | |" +
			"|k| | | | | | | |",
			&Move{from: 13, to: 5, promoteTo: Knight}, White},
		{"   |R| | | | | |R|K|" +
			"| |P| | | | |P|P|" +
			"| | |P|n| | | | |" +
			"|Q|P| | | |P| | |" +
			"| | | | | | | | |" +
			"| | | |p| | |p| |" +
			"|p|p|p| | | |b|p|" +
			"| |k|r| | | | |r|",
			&Move{from: 19, to: 13}, White},
		{"   | |R| | | | | | |" +
			"| |N| | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | |K| | |" +
			"| | | | | | | | |" +
			"| | | | | | | |k|",
			&Move{from: 45, to: 53}, Black},
	}

	for i, test := range tests {
		b := NewBoardFromDiagram(test.diagram)
		b.turn = test.turn
		b.nbMoves = 100
		move := ai.GetBestMove(b, 0, 2)

		if !move.Equals(test.move) {
			t.Errorf("test %d: expected move %s instead of %s",
				i, test.move, move)
		}
	}
}

func TestAIPruneNodes(t *testing.T) {
	ai := NewAI()

	tests := []struct {
		list         []PrunableNode
		turn         Color
		pruneRate    float64
		minKeptNodes uint
		scores       []float64
	}{
		{[]PrunableNode{
			PrunableNode{node: &Node{score: 1.5}},
			PrunableNode{node: &Node{score: 2.1}},
			PrunableNode{node: &Node{score: 1.7}},
			PrunableNode{node: &Node{score: 2.0}},
		}, Black, 0.5, 1, []float64{1.5, 1.7}},
		{[]PrunableNode{
			PrunableNode{node: &Node{score: 1.5}},
			PrunableNode{node: &Node{score: 2.1}},
			PrunableNode{node: &Node{score: 1.7}},
			PrunableNode{node: &Node{score: 2.0}},
		}, Black, 0.5, 3, []float64{1.5, 1.7, 2.0}},
		{[]PrunableNode{
			PrunableNode{node: &Node{score: 1.5}},
		}, Black, 0.5, 42, []float64{1.5}},
		{[]PrunableNode{
			PrunableNode{node: &Node{score: 1.5}},
			PrunableNode{node: &Node{score: 2.1}},
			PrunableNode{node: &Node{score: 1.7}},
			PrunableNode{node: &Node{score: 2.0}},
		}, White, 0.5, 2, []float64{2.0, 2.1}},
	}

	for i, test := range tests {
		list := ai.pruneNodes(test.list,
			test.turn, test.pruneRate, test.minKeptNodes)
		if len(list) != len(test.scores) {
			t.Fatalf("test %d: expected len of %d instead of %d",
				i, len(test.scores), len(list))
		}

		for j := range list {
			if test.scores[j] != list[j].node.score {
				t.Fatalf("test %d: expected elem %d "+
					"to be %.2f instead of %.2f",
					i, j, test.scores[j], list[j].node.score)
			}
		}
	}
}

func TestAIEvalCompare(t *testing.T) {
	tests := []struct {
		turn    Color
		better  string
		nbMoves int
		worse   string
	}{
		{
			White,
			"    | | | | | | | | |" +
				"| | | | | | |K| |" +
				"| | | | | | |P| |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
			100,
			"    | | | | | | | | |" +
				"| | | | | | |K| |" +
				"| | | | | | | | |" +
				"| | | | | | |P| |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
		},
		{
			White,
			"    | | | | | | |r| |" +
				"| | | | | |r| | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | |K|" +
				"|k| | | | | | | |",
			100,
			"    | | | | |r| | | |" +
				"| | | | | |r| | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | |K| |",
		},
		{
			White,
			"    |R|N|B|Q|K|B|N|R|" +
				"|P|P|P|P| |P|P|P|" +
				"| | | | | | | | |" +
				"| | | | |P| | | |" +
				"| | |b| |p| | | |" +
				"| | | | | | | | |" +
				"|p|p|p|p| |p|p|p|" +
				"|r|n|b|q|k| |n|r|",
			0,
			"    |R|N|B|Q|K|B|N|R|" +
				"|P|P|P|P| |P|P|P|" +
				"| | | | | | | | |" +
				"| |b| | |P| | | |" +
				"| | | | |p| | | |" +
				"| | | | | | | | |" +
				"|p|p|p|p| |p|p|p|" +
				"|r|n|b|q|k| |n|r|",
		},
		{
			White,
			"    |R| |B|Q|K|B|N|R|" +
				"|P|P|P|P|P|P|P|P|" +
				"|N| | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | |p| | | |" +
				"| | | | | | | | |" +
				"|p|p|p|p| |p|p|p|" +
				"|r|n|b|q|k|b|n|r|",
			0,
			"    |R| |B|Q|K|B|N|R|" +
				"|P|P|P|P|P|P|P|P|" +
				"| | |N| | | | | |" +
				"| | | | | | | | |" +
				"| | | | |p| | | |" +
				"| | | | | | | | |" +
				"|p|p|p|p| |p|p|p|" +
				"|r|n|b|q|k|b|n|r|",
		},
		{
			Black,
			"    |R|N|B|Q|K| |N|R|" +
				"|P|P|P| | |P|P|P|" +
				"| | | | | | | | |" +
				"| | | |P|P| | | |" +
				"| |p| | |p| | | |" +
				"| | |n| | | | | |" +
				"| |p|p|p| |p|p|p|" +
				"|r| |b|q|k|b|n|r|",
			0,
			"    |R|N|B|Q|K| |N|R|" +
				"|P|P|P| | |P|P|P|" +
				"| | | | | | | | |" +
				"| | | |P|P| | | |" +
				"| |B| | |p| | | |" +
				"|p| |n|p| | | | |" +
				"| |p|p| | |p|p|p|" +
				"|r| |b|q|k|b|n|r|",
		},
	}

	ai := NewAI()

	for i, test := range tests {
		better := NewBoardFromDiagram(test.better)
		worse := NewBoardFromDiagram(test.worse)
		better.turn = test.turn
		worse.turn = test.turn
		better.nbMoves = test.nbMoves
		worse.nbMoves = test.nbMoves
		betterScore := ai.evalPosition(better)
		worseScore := ai.evalPosition(worse)

		if worseScore >= betterScore {
			t.Errorf("test %d: the worse position has a better"+
				" or equal score (%.3f) than the best one (%.3f)",
				i, worseScore, betterScore)
		}
	}
}

func TestAIEvalEndgamePawn(t *testing.T) {
	ai := NewAI()
	b := NewBoardFromDiagram(
		"    | | | | | | | | |" +
			"| | | | | | |K| |" +
			"| | | | | | |P| |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"|k| | | | | | | |")

	score := ai.evalPosition(b)

	b = NewBoardFromDiagram(
		"    | | | | | | | | |" +
			"| | | | | | |K| |" +
			"| | | | | | | | |" +
			"| | | | | | |P| |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"|k| | | | | | | |")

	score2 := ai.evalPosition(b)

	if score <= score2 {
		t.Fatalf("second diagram should be better for black")
	}
}

func TestAIEvalPosition(t *testing.T) {
	ai := NewAI()
	b := NewBoard()

	score := ai.evalPosition(b)
	if score > 0.0001 || score < -0.0001 {
		t.Fatalf("expected score of 0 instead of %f", score)
	}
}

func TestAIBadMoves(t *testing.T) {
	if testing.Short() {
		return
	}

	ai := NewAI()

	tests := []struct {
		diagram string
		turn    Color
		badMove Move
	}{
		{"   |R| |B|Q|K|B|N|R|" +
			"|P|P|P| | |P|P|P|" +
			"| | |N| | | | | |" +
			"| | | |n|P| | | |" +
			"| | | | |p| | | |" +
			"| | | | | |n| | |" +
			"|p|p|p|p| |p|p|p|" +
			"|r| |b|q|k|b| |r|",
			Black,
			Move{from: 3, to: 27}},
		{"   |R|N|B|Q|K|B|N|R|" +
			"|P|P|P|P| |P|P|P|" +
			"| | | | | | | | |" +
			"| | | | |P| | | |" +
			"| | | | |p| | | |" +
			"| | |n| | | | | |" +
			"|p|p|p|p| |p|p|p|" +
			"|r| |b|q|k|b|n|r|",
			Black,
			Move{from: 11, to: 27}},
		{"   |R|N| |Q| |R|K| |" +
			"|P|P|P| | |P|P|P|" +
			"| | | |B| | | | |" +
			"| | | |N|P|B| | |" +
			"| | |b| | | | | |" +
			"| | |n|p| | | | |" +
			"|p|p|p| |n|p|p|p|" +
			"|r| |b|q| |r|k| |",
			Black,
			Move{from: 21, to: 27}},
		{"   |R| | | | |R| |K|" +
			"| |P|P| | |P|P| |" +
			"|P| | | | | | |P|" +
			"| | | | | |n| | |" +
			"| | |b|Q| | | | |" +
			"| |q| |p| | | | |" +
			"|p|p| | | |p|p| |" +
			"|r| | | |r| |k| |",
			Black,
			Move{from: 35, to: 34}},
		{"   |R| | |Q| |R|K| |" +
			"| |P|P| | |P|P| |" +
			"|P| | |B| | | |P|" +
			"| | | | |P| | |q|" +
			"| | | | |p| | | |" +
			"|p|p| |p| | |b| |" +
			"| |p| | | |p|p|p|" +
			"|r| | | | |r|k| |",
			Black,
			Move{from: 19, to: 40}},
		{"   |R| | | | |R|K| |" +
			"| |P|P| | |P|P| |" +
			"|P| | | | | | |P|" +
			"| | | | |b| | |q|" +
			"| | | | |Q| | | |" +
			"|p|p| | | | | | |" +
			"| | | | | |p|p|p|" +
			"|r| | | |r| |k| |",
			Black,
			Move{from: 36, to: 60}},
		{"   |R| | |Q| |R|K| |" +
			"|P|P|P| | |P|P|P|" +
			"| | |N| | | | | |" +
			"| | | |N|P| | | |" +
			"| | |b| | | | | |" +
			"|p| |p|p| |q| |p|" +
			"| | |p|b| |p|p| |" +
			"|r| | | |k| | |r|",
			Black,
			Move{from: 27, to: 42}},
	}

	for i, test := range tests {
		if i != 0 {
			continue
		}
		b := NewBoardFromDiagram(test.diagram)
		b.turn = test.turn

		move := ai.GetBestMove(b, 0, 2)
		if move.Equals(&test.badMove) {
			t.Errorf("test %d: ai played a bad move %v", i, move)
		}
	}
}

func TestAIGetBestMoveFromPos(t *testing.T) {
	ai := NewAI()
	ai.Genes["PruneRatio"] = 0.0
	b := NewBoardFromDiagram(
		"    |R|N|B|Q|K| |N|R|" +
			"|P|P|P|P| |P|P|P|" +
			"| | | | | | | | |" +
			"| | | | |P| | | |" +
			"| |B| | |p| | | |" +
			"|p| |n| | | | | |" +
			"| |p|p|p| |p|p|p|" +
			"|r| |b|q|k|b|n|r|")

	b.turn = Black
	ai.evalPosition(b)
}

func TestAIMutation(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	ai := NewAIEmpty()
	ai.mute(3, 1.0)
}
