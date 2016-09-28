package geneticchess

import (
	"fmt"
	"testing"
	"time"
)

func checkMoves(t *testing.T, moves []Move, expected []Move, errStr string) {
	expectedMap := make(map[string]Move)
	for _, move := range expected {
		key := fmt.Sprintf("%dx%d:%v", move.from, move.to, move.promoteTo)
		expectedMap[key] = move
	}

	movesMap := make(map[string]Move)
	for _, move := range moves {
		key := fmt.Sprintf("%dx%d:%v", move.from, move.to, move.promoteTo)
		movesMap[key] = move
	}

	for key, move := range expectedMap {
		_, ok := movesMap[key]
		if ok {
			continue
		}

		t.Errorf("%s: missing move %s", errStr, move.String())
	}

	for key, move := range movesMap {
		_, ok := expectedMap[key]
		if ok {
			continue
		}

		t.Errorf("%s: unexpected move %s", errStr, move.String())
	}
}

func TestBoardClone(t *testing.T) {
	b := NewBoard()
	b2 := b.clone()

	if b.hash() != b2.hash() {
		t.Fatalf("cloned board does not match original one")
	}

	_ = b.Move(&Move{from: 48, to: 40})
	_ = b2.Move(&Move{from: 48, to: 40})

	if b.hash() != b2.hash() {
		t.Fatalf("cloned board does not match original one after a move")
	}

	_ = b.GetMoves()

	if b.hash() != b2.hash() {
		t.Fatalf("cloned board does not match original one after get moves")
	}
}

func TestBoardNewBoardFromDiagram(t *testing.T) {
	diagram := "|R|N|B|Q|K|B|N|R|\n" +
		"|P|P|P|P|P|P|P|P|\n" +
		"| | | | | | | | |\n" +
		"| | | | | | | | |\n" +
		"| | | | | | | | |\n" +
		"| | | | | | | | |\n" +
		"|p|p|p|p|p|p|p|p|\n" +
		"|r|n|b|q|k|b|n|r|\n"
	b := NewBoardFromDiagram(diagram)

	if b.getDump() != diagram {
		t.Fatalf("expected %s instead of %s",
			diagram, b.hash())
	}
}

func TestBoardKnightMoves(t *testing.T) {
	b := NewEmptyBoard()

	b.squares[0] = &Piece{
		kind:  Knight,
		color: White,
	}
	m := b.getKnightMoves(0)
	checkMoves(t, m, []Move{
		Move{from: 0, to: 10, promoteTo: 0},
		Move{from: 0, to: 17, promoteTo: 0},
	}, "bad knight move 1")
	b.squares[0] = &Piece{}

	b.squares[7] = &Piece{
		kind:  Knight,
		color: White,
	}
	m = b.getKnightMoves(7)
	checkMoves(t, m, []Move{
		Move{from: 7, to: 13, promoteTo: 0},
		Move{from: 7, to: 22, promoteTo: 0},
	}, "bad knight move 2")
	b.squares[7] = &Piece{}

	b.squares[27] = &Piece{
		kind:  Knight,
		color: White,
	}
	m = b.getKnightMoves(27)
	checkMoves(t, m, []Move{
		Move{from: 27, to: 17, promoteTo: 0},
		Move{from: 27, to: 10, promoteTo: 0},
		Move{from: 27, to: 12, promoteTo: 0},
		Move{from: 27, to: 21, promoteTo: 0},
		Move{from: 27, to: 37, promoteTo: 0},
		Move{from: 27, to: 44, promoteTo: 0},
		Move{from: 27, to: 42, promoteTo: 0},
		Move{from: 27, to: 33, promoteTo: 0},
	}, "bad knight move 3")
	b.squares[27] = &Piece{}

	b.squares[56] = &Piece{
		kind:  Knight,
		color: White,
	}
	m = b.getKnightMoves(56)
	checkMoves(t, m, []Move{
		Move{from: 56, to: 41, promoteTo: 0},
		Move{from: 56, to: 50, promoteTo: 0},
	}, "bad knight move 4")
	b.squares[56] = &Piece{}

	b.squares[63] = &Piece{
		kind:  Knight,
		color: White,
	}
	m = b.getKnightMoves(63)
	checkMoves(t, m, []Move{
		Move{from: 63, to: 53, promoteTo: 0},
		Move{from: 63, to: 46, promoteTo: 0},
	}, "bad knight move 5")
	b.squares[63] = &Piece{}
}

func TestBoardRookMoves(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[27] = &Piece{
		kind:  Rook,
		color: White,
	}
	m := b.getRookMoves(27)

	checkMoves(t, m, []Move{
		Move{from: 27, to: 19, promoteTo: 0},
		Move{from: 27, to: 11, promoteTo: 0},
		Move{from: 27, to: 03, promoteTo: 0},

		Move{from: 27, to: 35, promoteTo: 0},
		Move{from: 27, to: 43, promoteTo: 0},
		Move{from: 27, to: 51, promoteTo: 0},
		Move{from: 27, to: 59, promoteTo: 0},

		Move{from: 27, to: 26, promoteTo: 0},
		Move{from: 27, to: 25, promoteTo: 0},
		Move{from: 27, to: 24, promoteTo: 0},

		Move{from: 27, to: 28, promoteTo: 0},
		Move{from: 27, to: 29, promoteTo: 0},
		Move{from: 27, to: 30, promoteTo: 0},
		Move{from: 27, to: 31, promoteTo: 0},
	}, "bad rook moves")
}

func TestBoardRookMovesWithFriendlyPiecesAround(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[27] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[26] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[19] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[28] = &Piece{
		kind:  Pawn,
		color: White,
	}
	m := b.getRookMoves(27)

	checkMoves(t, m, []Move{
		Move{from: 27, to: 35, promoteTo: 0},
		Move{from: 27, to: 43, promoteTo: 0},
		Move{from: 27, to: 51, promoteTo: 0},
		Move{from: 27, to: 59, promoteTo: 0},
	}, "bad rook moves")
}

func TestBoardRookMovesWithEnemyPiecesAround(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[27] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[26] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[19] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[28] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	m := b.getRookMoves(27)

	checkMoves(t, m, []Move{
		Move{from: 27, to: 19, promoteTo: 0},

		Move{from: 27, to: 35, promoteTo: 0},
		Move{from: 27, to: 43, promoteTo: 0},
		Move{from: 27, to: 51, promoteTo: 0},
		Move{from: 27, to: 59, promoteTo: 0},

		Move{from: 27, to: 26, promoteTo: 0},

		Move{from: 27, to: 28, promoteTo: 0},
	}, "bad rook moves")
}

func TestBoardBishopMoves(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[27] = &Piece{
		kind:  Bishop,
		color: White,
	}
	m := b.getBishopMoves(27)

	checkMoves(t, m, []Move{
		Move{from: 27, to: 18, promoteTo: 0},
		Move{from: 27, to: 9, promoteTo: 0},
		Move{from: 27, to: 0, promoteTo: 0},

		Move{from: 27, to: 20, promoteTo: 0},
		Move{from: 27, to: 13, promoteTo: 0},
		Move{from: 27, to: 6, promoteTo: 0},

		Move{from: 27, to: 36, promoteTo: 0},
		Move{from: 27, to: 45, promoteTo: 0},
		Move{from: 27, to: 54, promoteTo: 0},
		Move{from: 27, to: 63, promoteTo: 0},

		Move{from: 27, to: 34, promoteTo: 0},
		Move{from: 27, to: 41, promoteTo: 0},
		Move{from: 27, to: 48, promoteTo: 0},
	}, "bad bishop moves")
}

func TestBoardBishopMovesWithEnemyPieceAround(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[40] = &Piece{
		kind:  Bishop,
		color: White,
	}
	b.squares[26] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	m := b.getBishopMoves(40)

	checkMoves(t, m, []Move{
		Move{from: 40, to: 49, promoteTo: 0},
		Move{from: 40, to: 58, promoteTo: 0},

		Move{from: 40, to: 33, promoteTo: 0},
		Move{from: 40, to: 26, promoteTo: 0},
	}, "bad bishop moves")
}

func TestBoardQueenMoves(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[9] = &Piece{
		kind:  Queen,
		color: White,
	}
	m := b.getQueenMoves(9)

	checkMoves(t, m, []Move{
		Move{from: 9, to: 0, promoteTo: 0},
		Move{from: 9, to: 18, promoteTo: 0},
		Move{from: 9, to: 27, promoteTo: 0},
		Move{from: 9, to: 36, promoteTo: 0},
		Move{from: 9, to: 45, promoteTo: 0},
		Move{from: 9, to: 54, promoteTo: 0},
		Move{from: 9, to: 63, promoteTo: 0},

		Move{from: 9, to: 16, promoteTo: 0},
		Move{from: 9, to: 2, promoteTo: 0},

		Move{from: 9, to: 8, promoteTo: 0},
		Move{from: 9, to: 10, promoteTo: 0},
		Move{from: 9, to: 11, promoteTo: 0},
		Move{from: 9, to: 12, promoteTo: 0},
		Move{from: 9, to: 13, promoteTo: 0},
		Move{from: 9, to: 14, promoteTo: 0},
		Move{from: 9, to: 15, promoteTo: 0},

		Move{from: 9, to: 1, promoteTo: 0},
		Move{from: 9, to: 17, promoteTo: 0},
		Move{from: 9, to: 25, promoteTo: 0},
		Move{from: 9, to: 33, promoteTo: 0},
		Move{from: 9, to: 41, promoteTo: 0},
		Move{from: 9, to: 49, promoteTo: 0},
		Move{from: 9, to: 57, promoteTo: 0},
	}, "bad queen moves")
}

func TestBoardKingMoves(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[35] = &Piece{
		kind:  King,
		color: White,
	}
	m := b.getKingMoves(35)

	checkMoves(t, m, []Move{
		Move{from: 35, to: 34, promoteTo: 0},
		Move{from: 35, to: 26, promoteTo: 0},
		Move{from: 35, to: 27, promoteTo: 0},
		Move{from: 35, to: 28, promoteTo: 0},
		Move{from: 35, to: 36, promoteTo: 0},
		Move{from: 35, to: 44, promoteTo: 0},
		Move{from: 35, to: 43, promoteTo: 0},
		Move{from: 35, to: 42, promoteTo: 0},
	}, "bad king moves")
}

func TestBoardPawnMoves(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[27] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[36] = &Piece{
		kind:  Pawn,
		color: White,
	}
	m := b.getPawnMoves(36)

	checkMoves(t, m, []Move{
		Move{from: 36, to: 28, promoteTo: 0},
		Move{from: 36, to: 27, promoteTo: 0},
	}, "bad pawn moves")
}

func TestBoardPawnPromotion(t *testing.T) {
	b := NewEmptyBoard()
	b.squares[5] = &Piece{
		kind:  Rook,
		color: Black,
	}
	b.squares[12] = &Piece{
		kind:  Pawn,
		color: White,
	}
	m := b.getPawnMoves(12)

	checkMoves(t, m, []Move{
		Move{from: 12, to: 4, promoteTo: Queen},
		Move{from: 12, to: 4, promoteTo: Rook},
		Move{from: 12, to: 4, promoteTo: Bishop},
		Move{from: 12, to: 4, promoteTo: Knight},
		Move{from: 12, to: 5, promoteTo: Queen},
		Move{from: 12, to: 5, promoteTo: Rook},
		Move{from: 12, to: 5, promoteTo: Bishop},
		Move{from: 12, to: 5, promoteTo: Knight},
	}, "bad pawn moves")

	b = NewEmptyBoard()
	b.squares[63] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[54] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	m = b.getPawnMoves(54)

	checkMoves(t, m, []Move{
		Move{from: 54, to: 62, promoteTo: Queen},
		Move{from: 54, to: 62, promoteTo: Rook},
		Move{from: 54, to: 62, promoteTo: Bishop},
		Move{from: 54, to: 62, promoteTo: Knight},
		Move{from: 54, to: 63, promoteTo: Queen},
		Move{from: 54, to: 63, promoteTo: Rook},
		Move{from: 54, to: 63, promoteTo: Bishop},
		Move{from: 54, to: 63, promoteTo: Knight},
	}, "bad pawn moves")
}

func TestBoardPawnEnPassant(t *testing.T) {
	tests := []struct {
		diagram   string
		turn      Color
		firstMove Move
		enPassant Move
	}{
		{
			"    | | | | | | | |K|" +
				"| | | | | | |P| |" +
				"| | | | | | | | |" +
				"| | | | | |p| | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | |p| |" +
				"| | | | |r| |k| |",
			Black,
			Move{from: 14, to: 30},
			Move{from: 29, to: 22},
		},
		{
			"    | | | | | | | |K|" +
				"| | | | | | |P| |" +
				"| | | | | | | | |" +
				"| | | | | | | |p|" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | |p| |" +
				"| | | | |r| |k| |",
			Black,
			Move{from: 14, to: 30},
			Move{from: 31, to: 22},
		},
		{
			"    | | | | | | | |K|" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | |P| | | | | |" +
				"| | | | | | | | |" +
				"| | | |p| | | | |" +
				"| | | | |r| |k| |",
			White,
			Move{from: 51, to: 35},
			Move{from: 34, to: 43},
		},
		{
			"    | | | | | | | |K|" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| |P|P| | | | | |" +
				"| | | | | | | | |" +
				"|p| | | | | | | |" +
				"| | | | |r| |k| |",
			White,
			Move{from: 48, to: 32},
			Move{from: 33, to: 40},
		},
	}

	for i, test := range tests {
		var enPassant *Move
		b := NewBoardFromDiagram(test.diagram)
		b.turn = test.turn

		_ = b.Move(&test.firstMove)
		m := b.GetMoves()

		for _, move := range m {
			if move.Equals(&test.enPassant) {
				enPassant = &move
				break
			}
		}

		if enPassant == nil {
			t.Fatalf("test %d: cannot find en passant move %v",
				i, test.enPassant)
		}

		_ = b.Move(enPassant)

		if b.squares[test.firstMove.to] != nil {
			t.Errorf("test %d: pawn at %d has not been taken",
				i, test.firstMove.to)
		}
	}
}

func TestBoardGameMoves(t *testing.T) {
	b := NewBoard()
	m := b.GetMoves()

	if len(m) != 20 {
		t.Fatalf("white should have 20 moves instead of %d", len(m))
	}

	b.Move(&m[0])

	if b.squares[48] != nil {
		t.Fatalf("square 48 should be empty")
	}
	if b.squares[49] == nil {
		t.Fatalf("square 49 should not be empty")
	}
}

func TestBoardGameRandomMoves(t *testing.T) {
	b := NewBoard()

	for {
		m := b.GetMoves()

		if len(m) == 0 {
			break
		}

		state := b.Move(&m[0])
		switch state {
		case StateDrawByRepetition:
			return
		}
	}
}

func TestBoardDrawByRepetition(t *testing.T) {
	b := NewBoard()

	state := b.Move(&Move{from: 57, to: 40})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 1, to: 16})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 40, to: 57})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 16, to: 1})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 57, to: 40})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 1, to: 16})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}

	state = b.Move(&Move{from: 40, to: 57})
	if state != StatePlaying {
		t.Fatalf("expected state to be playing instead of %s", state.String())
	}
	state = b.Move(&Move{from: 16, to: 1})

	if state != StateDrawByRepetition {
		t.Fatalf("expected state to be draw by repetition instead of %s",
			state.String())
	}
}

func TestBoardDrawByStalemate(t *testing.T) {
	b := NewEmptyBoard()

	b.squares[0] = &Piece{
		kind:  King,
		color: Black,
	}
	b.squares[16] = &Piece{
		kind:  King,
		color: White,
	}
	b.squares[17] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[18] = &Piece{
		kind:  Rook,
		color: White,
	}

	state := b.Move(&Move{from: 18, to: 10})
	if state != StateDrawByStalemate {
		t.Fatalf("expected state to be draw by stalemate instead of %s",
			state.String())
	}
}

func TestBoardCheck(t *testing.T) {
	b := NewBoard()

	check := b.isCheck()
	if check == true {
		t.Fatalf("position should not give a check")
	}

	_ = b.Move(&Move{from: 52, to: 44})

	check = b.isCheck()
	if check == true {
		t.Fatalf("position should not give a check")
	}

	_ = b.Move(&Move{from: 11, to: 19})

	check = b.isCheck()
	if check == true {
		t.Fatalf("position should not give a check")
	}

	_ = b.Move(&Move{from: 61, to: 25})

	check = b.isCheck()
	if check != true {
		t.Fatalf("position should give a check")
	}
}

func TestBoardCheckmate(t *testing.T) {
	b := NewBoard()
	_ = b.Move(&Move{from: 52, to: 36})
	_ = b.Move(&Move{from: 12, to: 28})
	_ = b.Move(&Move{from: 59, to: 45})
	_ = b.Move(&Move{from: 8, to: 16})
	_ = b.Move(&Move{from: 61, to: 34})
	_ = b.Move(&Move{from: 16, to: 24})

	state := b.Move(&Move{from: 45, to: 13})
	if state != StateWhiteWins {
		t.Fatalf("expected %s instead of %s", StateWhiteWins, state)
	}
}

func TestBoardCastlingPossible(t *testing.T) {
	tests := []struct {
		diagram string
		from    Position
		king    Position
		rook    Position
		turn    Color
	}{
		{
			diagram: "|R| |B|Q|K|B|N|R|\n" +
				"|P|P| | |P|P|P|P|\n" +
				"| | |N| | | | | |\n" +
				"| | |P|P| | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | |n|p| |\n" +
				"|p|p|p|p|p|p|b|p|\n" +
				"|r|n|b|q|k| | |r|\n",
			from: 60,
			king: 62,
			rook: 61,
			turn: White,
		},
		{
			diagram: "|R| |B|Q|K| | |R|\n" +
				"|P|P| | |P|P|P|P|\n" +
				"| | |N| | | | | |\n" +
				"| | |P|P| | | | |\n" +
				"| | | | | | | | |\n" +
				"| | |p| | |n|p| |\n" +
				"|p|p| |p|p|p|b|p|\n" +
				"|r|n|b|q|k| | |r|\n",
			from: 4,
			king: 6,
			rook: 5,
			turn: Black,
		},
		{
			diagram: "|R| |B|Q|K|B| |R|\n" +
				"|P|P| | |P|P|P|P|\n" +
				"| | |N| | | |N| |\n" +
				"| | |P|P| | | | |\n" +
				"| | | | | | | | |\n" +
				"| |p|n| |p| | | |\n" +
				"|p|b|p|p|q|p|p|p|\n" +
				"|r| | | |k|b|n|r|\n",
			from: 60,
			king: 58,
			rook: 59,
			turn: White,
		},
		{
			diagram: "|R| | | |K|B|N|R|\n" +
				"|P|P| | |P|P|P|P|\n" +
				"| | |N| | | | | |\n" +
				"| | |P|P| | | | |\n" +
				"| | | | | | | | |\n" +
				"| |p|n| |p| | | |\n" +
				"|p|b|p|p| |p|p|p|\n" +
				"|r| | |q|k|b|n|r|\n",
			from: 4,
			king: 2,
			rook: 3,
			turn: Black,
		},
	}

	for i, test := range tests {
		b := NewBoardFromDiagram(test.diagram)
		b.turn = test.turn
		m := b.GetMoves()

		var castle *Move

		for _, move := range m {
			if move.from == test.from && move.to == test.king {
				castle = &move
				break
			}
		}

		if castle == nil {
			t.Fatalf("test %d: cannot castle", i)
		}

		_ = b.Move(castle)

		if b.squares[test.king] == nil ||
			b.squares[test.king].kind != King {
			t.Fatalf("test %d: castling failed: king expected on %d",
				i, test.king)
		}
		if b.squares[test.rook] == nil ||
			b.squares[test.rook].kind != Rook {
			t.Fatalf("test %d: castling failed: rook expected on %d",
				i, test.rook)
		}
	}
}

func TestBoardCastlingImpossible(t *testing.T) {
	tests := []struct {
		diagram string
		from    Position
		to      Position
		turn    Color
	}{
		{
			diagram: "|R| |B|Q|K|B|N|R|\n" +
				"|P|P| | |P|P|P|P|\n" +
				"| | | | | | | | |\n" +
				"| | |P|P| | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | |N| |n|p| |\n" +
				"|p|p|p|p|p|p|b|p|\n" +
				"|r|n|b|q|k| | |r|\n",
			from: 60,
			to:   62,
			turn: White,
		},
		{
			diagram: "|K| | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | |P| | |\n" +
				"|r| | | |k| | | |\n",
			from: 60,
			to:   58,
			turn: White,
		},
		{
			diagram: "|K| | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | |R| | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"|r| | | |k| | | |\n",
			from: 60,
			to:   58,
			turn: White,
		},
		{
			diagram: "| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | |K| | | | | |\n" +
				"|r| | | |k| | | |\n",
			from: 60,
			to:   58,
			turn: White,
		},
		{
			diagram: "| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | | |\n" +
				"| | | | | | | |N|\n" +
				"| | | | | | | | |\n" +
				"| | | | |k| | |r|\n",
			from: 60,
			to:   62,
			turn: White,
		},
	}

	for i, test := range tests {
		b := NewBoardFromDiagram(test.diagram)
		b.turn = test.turn
		m := b.GetMoves()

		var castle *Move

		for _, move := range m {
			if move.from == test.from && move.to == test.to {
				castle = &move
				break
			}
		}

		if castle != nil {
			t.Fatalf("test %d: can castle when not allowed", i)
		}
	}
}

func TestBoardRand(t *testing.T) {
	b := NewBoard()

	for i := int(time.Now().Second()); ; i++ {
		m := b.GetMoves()
		state := b.Move(&m[i%len(m)])
		if state != StatePlaying {
			break
		}
	}
}

func TestBoardGetMovesBug(t *testing.T) {
	s :=
		"    |R|N|N| |K| | |R|" +
			"|P| | | | |P| |P|" +
			"| | | | | | |P| |" +
			"| |P| | |p| | | |" +
			"| |p| | | | | | |" +
			"| |p| | | | | |B|" +
			"|p| | | | | | |p|" +
			"| | | | |k| | |r|"
	b := NewBoardFromDiagram(s)
	b.turn = Black

	_ = b.GetMoves()
}

func TestBoardPromotionBug(t *testing.T) {
	s :=
		"    | | | | |K| | | |" +
			"|p| | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"| | | | | | | | |" +
			"|k| | | | | | | |"
	b := NewBoardFromDiagram(s)

	b.GetMoves()
}

func TestBoardBoardHasEnoughMaterial(t *testing.T) {
	tests := []struct {
		diagram           string
		hasEnoughMaterial bool
	}{
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
			false},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | |n| | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
			false},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | |N| |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
			false},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | |B|" +
				"| | | | | | | | |" +
				"|k| | | | | | | |",
			false},
		{
			"    |B| | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|k| | | | | | |b|",
			false},
		{
			"    | |B| |b|K|B| |b|" +
				"|B| |b| |B| |b| |" +
				"| |b| |B| |b| |B|" +
				"|b| |B| |b| |B| |" +
				"| |B| |b| |B| |b|" +
				"|B| |b| |B| |b| |" +
				"| |b| |B| |b| |B|" +
				"|b|K|B| |b| |B| |",
			false},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|q| | | | | | | |" +
				"|k| | | | | | | |",
			true},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | |R| |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|n| | | | | | | |" +
				"|k| | | | | | | |",
			true},
		{
			"    | | | | |K| | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"| | | | | | | | |" +
				"|p| | | | | | | |" +
				"|k| | | | | | | |",
			true},
	}

	for i, test := range tests {
		b := NewBoardFromDiagram(test.diagram)
		res := b.hasEnoughMaterial()
		if res != test.hasEnoughMaterial {
			str := ""

			if test.hasEnoughMaterial == false {
				str = "not "
			}

			t.Errorf("test %d: should %sbe able to continue",
				i, str)
		}
	}
}
