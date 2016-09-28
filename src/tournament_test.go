package geneticchess

import (
	"testing"
	"time"
)

func getAI() *AI {
	return NewAI()
}

func TestTournament(t *testing.T) {
	if testing.Short() {
		return
	}

	tournament := NewTournament([]*AI{
		NewAI(), NewAI(),
	}, 1, 1, 2, 2, 0.5)

	res := tournament.Play(time.Millisecond, 0, 2, false)
	if len(res) == 0 {
		t.Fatalf("tournament results are empty")
	}
}
