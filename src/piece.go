package geneticchess

import (
	"log"
	"strings"
)

type PieceType uint8

type Piece struct {
	kind  PieceType
	color Color
	flags uint8
}

var pieceChars = map[PieceType]string{
	King:   "k",
	Queen:  "q",
	Rook:   "r",
	Knight: "n",
	Bishop: "b",
	Pawn:   "p",
}

var fullNames = map[PieceType]string{
	King:   "King",
	Queen:  "Queen",
	Rook:   "Rook",
	Knight: "Knight",
	Bishop: "Bishop",
	Pawn:   "Pawn",
}

const (
	Empty  PieceType = 0
	King   PieceType = 1
	Queen  PieceType = 2
	Rook   PieceType = 3
	Knight PieceType = 4
	Bishop PieceType = 5
	Pawn   PieceType = 6
)

const (
	HasMoved            uint8 = 1 << 0
	HasMovedRightBefore uint8 = 1 << 1
)

func (pt *PieceType) String() string {
	res, ok := pieceChars[*pt]
	if ok {
		return res
	}

	return ""
}

func (pt *PieceType) GetName() string {
	res, ok := fullNames[*pt]
	if ok {
		return res
	}

	return ""
}

func (p *Piece) String() string {
	c, ok := pieceChars[p.kind]
	if !ok {
		log.Panicf("unknown piece type: %d", p.kind)
	}

	if p.color == Black {
		return strings.ToUpper(c)
	}

	return c
}

func (p *Piece) clone() *Piece {
	return &Piece{
		kind:  p.kind,
		color: p.color,
		flags: p.flags,
	}
}

func StringToPieceType(str string) PieceType {
	switch str {
	case "k":
		return King
	case "q":
		return Queen
	case "r":
		return Rook
	case "b":
		return Bishop
	case "p":
		return Pawn
	default:
		return Empty
	}
}
