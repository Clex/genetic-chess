package geneticchess

import ()

type State int8

const (
	StatePlaying                    State = 0
	StateWhiteWins                  State = 1
	StateBlackWins                  State = 2
	StateDrawByRepetition           State = 3
	StateDrawByStalemate            State = 4
	StateDrawByInsufficientMaterial State = 5
)

func (s State) String() string {
	switch s {
	case StatePlaying:
		return "playing"
	case StateWhiteWins:
		return "white wins"
	case StateBlackWins:
		return "black wins"
	case StateDrawByRepetition:
		return "draw by repetition"
	case StateDrawByStalemate:
		return "draw by stalemate"
	case StateDrawByInsufficientMaterial:
		return "draw by insufficient material"
	default:
		panic("unknown state")
	}
}
