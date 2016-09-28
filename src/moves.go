package geneticchess

import (
	"fmt"
)

type Move struct {
	from      Position
	to        Position
	promoteTo PieceType
}

type Moves []Move

func (m *Moves) Append(board *Board, from Position, to Position) {
	move := Move{
		from: from,
		to:   to,
	}

	*m = append(*m, move)
}

func (m *Moves) AppendPromotion(board *Board,
	from Position, to Position, promotion PieceType) {
	move := Move{
		from:      from,
		to:        to,
		promoteTo: promotion,
	}

	*m = append(*m, move)
}

func (m *Moves) appendPawnMove(b *Board, from Position, to Position) {
	if to.isTopBorder() || to.isBottomBorder() {
		m.AppendPromotion(b, from, to, Queen)
		m.AppendPromotion(b, from, to, Rook)
		m.AppendPromotion(b, from, to, Knight)
		m.AppendPromotion(b, from, to, Bishop)
	} else {
		m.Append(b, from, to)
	}
}

func (m *Move) String() string {
	var prom string

	if m.promoteTo != Empty {
		prom = fmt.Sprintf("(%s)", m.promoteTo.String())
	}

	return fmt.Sprintf("%d:%d%s", m.from, m.to, prom)
}

func (m *Move) Equals(m2 *Move) bool {
	return m.from == m2.from &&
		m.to == m2.to &&
		m.promoteTo == m2.promoteTo
}
