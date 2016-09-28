package geneticchess

import (
	"testing"
)

func TestMoveEquals(t *testing.T) {
	m1 := Move{
		from:      0,
		to:        1,
		promoteTo: 0,
	}
	m2 := Move{
		from:      42,
		to:        43,
		promoteTo: 0,
	}
	m3 := Move{
		from:      7,
		to:        55,
		promoteTo: 0,
	}

	if m1.Equals(&m2) == true {
		t.Fatalf("m1 and m2 should not be equal")
	}
	if m1.Equals(&m3) == true {
		t.Fatalf("m1 and m3 should not be equal")
	}
	if m2.Equals(&m3) == true {
		t.Fatalf("m2 and m3 should not be equal")
	}
	if m1.Equals(&m1) == false {
		t.Fatalf("m1 and m1 should be equal")
	}
	if m2.Equals(&m2) == false {
		t.Fatalf("m2 and m2 should be equal")
	}
	if m3.Equals(&m3) == false {
		t.Fatalf("m3 and m3 should be equal")
	}
}

func TestMoveAppendLegal(t *testing.T) {
	var moves Moves

	b := NewBoard()

	moves.Append(b, 48, 40)

	if len(moves) != 1 {
		t.Fatalf("bad len: %d", len(moves))
	}
}
