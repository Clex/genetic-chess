package geneticchess

import (
	"testing"
)

func TestColorSwap(t *testing.T) {
	var c Color

	c = White
	c.swap()

	if c != Black {
		t.Fatalf("swap failed")
	}

	c.swap()

	if c != White {
		t.Fatalf("swap failed")
	}
}
