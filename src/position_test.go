package geneticchess

import (
	"testing"
)

func TestPositionGetRow(t *testing.T) {
	var pos Position

	pos = 0
	row := pos.getRow()
	if row != 0 {
		t.Fatalf("expected %d got %d", 0, row)
	}

	pos = 7
	row = pos.getRow()
	if row != 0 {
		t.Fatalf("expected %d got %d", 0, row)
	}

	pos = 8
	row = pos.getRow()
	if row != 1 {
		t.Fatalf("expected %d got %d", 1, row)
	}

	pos = 51
	row = pos.getRow()
	if row != 6 {
		t.Fatalf("expected %d got %d", 6, row)
	}

	pos = 63
	row = pos.getRow()
	if row != 7 {
		t.Fatalf("expected %d got %d", 7, row)
	}
}

func TestPositionGetCol(t *testing.T) {
	var pos Position

	pos = 0
	col := pos.getCol()
	if col != 0 {
		t.Fatalf("expected %d got %d", 0, col)
	}

	pos = 1
	col = pos.getCol()
	if col != 1 {
		t.Fatalf("expected %d got %d", 1, col)
	}

	pos = 52
	col = pos.getCol()
	if col != 4 {
		t.Fatalf("expected %d got %d", 4, col)
	}

	pos = 63
	col = pos.getCol()
	if col != 7 {
		t.Fatalf("expected %d got %d", 7, col)
	}
}

func TestPositionGetSquaresFromTop(t *testing.T) {
	var pos Position

	pos = 0
	squares := pos.getSquaresFromTop()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}

	pos = 7
	squares = pos.getSquaresFromTop()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}

	pos = 8
	squares = pos.getSquaresFromTop()
	if squares != 1 {
		t.Fatalf("expected %d got %d", 1, squares)
	}

	pos = 38
	squares = pos.getSquaresFromTop()
	if squares != 4 {
		t.Fatalf("expected %d got %d", 4, squares)
	}

	pos = 62
	squares = pos.getSquaresFromTop()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}
}

func TestPositionGetSquaresFromBottom(t *testing.T) {
	var pos Position

	pos = 0
	squares := pos.getSquaresFromBottom()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}

	pos = 7
	squares = pos.getSquaresFromBottom()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}

	pos = 8
	squares = pos.getSquaresFromBottom()
	if squares != 6 {
		t.Fatalf("expected %d got %d", 6, squares)
	}

	pos = 38
	squares = pos.getSquaresFromBottom()
	if squares != 3 {
		t.Fatalf("expected %d got %d", 3, squares)
	}

	pos = 62
	squares = pos.getSquaresFromBottom()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}
}

func TestPositionGetSquaresFromLeft(t *testing.T) {
	var pos Position

	pos = 0
	squares := pos.getSquaresFromLeft()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}

	pos = 7
	squares = pos.getSquaresFromLeft()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}

	pos = 8
	squares = pos.getSquaresFromLeft()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}

	pos = 38
	squares = pos.getSquaresFromLeft()
	if squares != 6 {
		t.Fatalf("expected %d got %d", 6, squares)
	}

	pos = 57
	squares = pos.getSquaresFromLeft()
	if squares != 1 {
		t.Fatalf("expected %d got %d", 1, squares)
	}
}

func TestPositionGetSquaresFromRight(t *testing.T) {
	var pos Position

	pos = 0
	squares := pos.getSquaresFromRight()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}

	pos = 7
	squares = pos.getSquaresFromRight()
	if squares != 0 {
		t.Fatalf("expected %d got %d", 0, squares)
	}

	pos = 8
	squares = pos.getSquaresFromRight()
	if squares != 7 {
		t.Fatalf("expected %d got %d", 7, squares)
	}

	pos = 38
	squares = pos.getSquaresFromRight()
	if squares != 1 {
		t.Fatalf("expected %d got %d", 1, squares)
	}

	pos = 58
	squares = pos.getSquaresFromRight()
	if squares != 5 {
		t.Fatalf("expected %d got %d", 5, squares)
	}
}
