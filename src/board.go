package geneticchess

import (
	"fmt"
	"log"
	"strings"
)

type Board struct {
	history map[string]int
	squares [64]*Piece
	turn    Color
	nbMoves int

	movesCache     Moves
	whiteKingCache Position
	blackKingCache Position
}

func NewBoard() *Board {
	turn := White
	board := &Board{
		turn:    turn,
		history: make(map[string]int),
	}

	board.initPieces()
	board.setKingCache()

	h := board.hash()
	board.history[h] = 1

	return board
}

func NewEmptyBoard() *Board {
	turn := White
	b := &Board{
		turn:    turn,
		history: make(map[string]int),
	}

	return b
}

func NewBoardFromDiagram(diag string) *Board {
	b := NewEmptyBoard()

	diag = strings.TrimSpace(diag)
	var pos Position

	for _, r := range diag {
		switch r {
		case ' ':
			b.squares[pos] = nil
		case 'k':
			b.squares[pos] = &Piece{kind: King, color: White}
		case 'K':
			b.squares[pos] = &Piece{kind: King, color: Black}
		case 'q':
			b.squares[pos] = &Piece{kind: Queen, color: White}
		case 'Q':
			b.squares[pos] = &Piece{kind: Queen, color: Black}
		case 'r':
			b.squares[pos] = &Piece{kind: Rook, color: White}
		case 'R':
			b.squares[pos] = &Piece{kind: Rook, color: Black}
		case 'b':
			b.squares[pos] = &Piece{kind: Bishop, color: White}
		case 'B':
			b.squares[pos] = &Piece{kind: Bishop, color: Black}
		case 'n':
			b.squares[pos] = &Piece{kind: Knight, color: White}
		case 'N':
			b.squares[pos] = &Piece{kind: Knight, color: Black}
		case 'p':
			b.squares[pos] = &Piece{kind: Pawn, color: White}
		case 'P':
			b.squares[pos] = &Piece{kind: Pawn, color: Black}
		default:
			continue
		}

		pos++
	}

	if pos != 64 {
		panic("bad board")
	}

	b.setKingCache()

	return b
}

func (b *Board) clone() *Board {
	newBoard := &Board{
		turn:           b.turn,
		nbMoves:        b.nbMoves,
		whiteKingCache: b.whiteKingCache,
		blackKingCache: b.blackKingCache,
	}

	newBoard.history = make(map[string]int)
	for key, val := range b.history {
		newBoard.history[key] = val
	}

	for key, val := range b.squares {
		if val == nil {
			continue
		}

		newBoard.squares[key] = val.clone()
	}

	return newBoard
}

func (b *Board) initPieces() {
	// Black pieces {{{
	b.squares[0] = &Piece{
		kind:  Rook,
		color: Black,
	}
	b.squares[1] = &Piece{
		kind:  Knight,
		color: Black,
	}
	b.squares[2] = &Piece{
		kind:  Bishop,
		color: Black,
	}
	b.squares[3] = &Piece{
		kind:  Queen,
		color: Black,
	}
	b.squares[4] = &Piece{
		kind:  King,
		color: Black,
	}
	b.squares[5] = &Piece{
		kind:  Bishop,
		color: Black,
	}
	b.squares[6] = &Piece{
		kind:  Knight,
		color: Black,
	}
	b.squares[7] = &Piece{
		kind:  Rook,
		color: Black,
	}
	b.squares[8] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[9] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[10] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[11] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[12] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[13] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[14] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	b.squares[15] = &Piece{
		kind:  Pawn,
		color: Black,
	}
	// }}}
	// White pieces {{{
	b.squares[56] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[57] = &Piece{
		kind:  Knight,
		color: White,
	}
	b.squares[58] = &Piece{
		kind:  Bishop,
		color: White,
	}
	b.squares[59] = &Piece{
		kind:  Queen,
		color: White,
	}
	b.squares[60] = &Piece{
		kind:  King,
		color: White,
	}
	b.squares[61] = &Piece{
		kind:  Bishop,
		color: White,
	}
	b.squares[62] = &Piece{
		kind:  Knight,
		color: White,
	}
	b.squares[63] = &Piece{
		kind:  Rook,
		color: White,
	}
	b.squares[48] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[49] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[50] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[51] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[52] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[53] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[54] = &Piece{
		kind:  Pawn,
		color: White,
	}
	b.squares[55] = &Piece{
		kind:  Pawn,
		color: White,
	}
	// }}}
}

func (b *Board) getDump() string {
	buf := ""

	for i, piece := range b.squares {
		val := " "

		if piece != nil {
			val = piece.String()
		}

		buf = buf + "|" + val

		if i%8 == 7 {
			buf = buf + "|\n"
		}
	}

	return buf
}

func (b *Board) Dump() {
	fmt.Println(b.getDump())
}

func (b *Board) hash() string {
	buf := ""

	for _, piece := range b.squares {
		if piece == nil {
			buf += " "
		} else {
			buf += piece.String()
		}
	}

	return buf
}

func (b *Board) hasEnoughMaterial() bool {
	lightBishops := 0
	darkBishops := 0
	knights := 0

	for i, sq := range b.squares {
		if sq == nil {
			continue
		}

		pos := Position(i)

		switch sq.kind {
		case Pawn:
			return true

		case Knight:
			knights++
			if knights > 1 {
				return true
			}

		case Bishop:
			if (pos.getRow()+pos.getCol())%2 == 0 {
				lightBishops++
			} else {
				darkBishops++
			}
			if lightBishops > 0 && darkBishops > 0 {
				return true
			}

		case Rook:
			return true

		case Queen:
			return true
		}
	}

	return false
}

func (b *Board) setKingCache() {
	for pos, val := range b.squares {
		if val == nil {
			continue
		}

		if val.kind == King {
			p := Position(pos)

			if val.color == White {
				b.whiteKingCache = p
			} else {
				b.blackKingCache = p
			}
		}
	}
}

func (b *Board) getKingPosition(color Color) Position {
	if color == White {
		if b.squares[b.whiteKingCache] == nil {
			panic("bad white king cache")
		}
		return b.whiteKingCache
	}
	return b.blackKingCache
}

func (b *Board) isKingTakable(move *Move) bool {
	tmpBoard := b.clone()
	tmpBoard.moveNoCheck(move)

	moves := tmpBoard.getMovesOpts(false, true, false)
	pos := tmpBoard.getKingPosition(!tmpBoard.turn)
	for _, move := range moves {
		if move.to == pos {
			return true
		}
	}

	return false
}

func (b *Board) areSquaresAttaqued(sq []Position) bool {
	board := b.clone()
	board.turn.swap()

	attackedSquares := make(map[Position]bool)

	moves := board.getMovesOpts(false, false, false)
	for _, move := range moves {
		attackedSquares[move.to] = true
	}

	for _, square := range sq {
		_, ok := attackedSquares[square]
		if ok == true {
			return true
		}
	}

	return false
}

func (b *Board) isCheck() bool {
	board := b.clone()
	board.turn.swap()

	moves := board.getMovesOpts(false, false, false)
	for _, move := range moves {
		p := board.squares[move.to]

		if p == nil {
			continue
		}

		if p.kind == King && p.color != board.turn {
			return true
		}
	}

	return false
}

func (b *Board) isMoveLegal(move *Move) bool {
	piece := b.squares[move.to]
	if piece != nil && piece.color == b.turn {
		return false
	}

	return !b.isKingTakable(move)
}

func (b *Board) getPawnMoves(pos Position) []Move {
	var moves Moves

	row := pos.getRow()
	isRightBorder := pos.isRightBorder()
	isLeftBorder := pos.isLeftBorder()

	switch b.squares[pos].color {
	case White:
		sq := b.squares[pos-8]
		if sq == nil {
			// Forward
			moves.appendPawnMove(b, pos, pos-8)

			if row == 6 {
				sq = b.squares[pos-16]
				if sq == nil {
					moves.appendPawnMove(b, pos, pos-16)
				}
			}
		}

		if !isRightBorder {
			// Right take
			sq := b.squares[pos-7]
			if sq != nil && sq.color == Black {
				moves.appendPawnMove(b, pos, pos-7)
			}
		}

		if !isLeftBorder {
			// Left take
			sq := b.squares[pos-9]
			if sq != nil && sq.color == Black {
				moves.appendPawnMove(b, pos, pos-9)
			}
		}

		if row == 3 {
			// En passant
			if !isRightBorder {
				rp := b.squares[pos+1]

				if rp != nil && rp.color == Black && rp.kind == Pawn &&
					rp.flags&HasMovedRightBefore != 0 {
					moves.appendPawnMove(b, pos, pos-7)
				}
			}
			if !isLeftBorder {
				rp := b.squares[pos-1]

				if rp != nil && rp.color == Black && rp.kind == Pawn &&
					rp.flags&HasMovedRightBefore != 0 {
					moves.appendPawnMove(b, pos, pos-9)
				}
			}
		}

	case Black:
		sq := b.squares[pos+8]
		if sq == nil {
			// Forward
			moves.appendPawnMove(b, pos, pos+8)

			if row == 1 {
				sq = b.squares[pos+16]
				if sq == nil {
					moves.appendPawnMove(b, pos, pos+16)
				}
			}
		}

		if !isLeftBorder {
			// Left take
			sq := b.squares[pos+7]
			if sq != nil && sq.color == White {
				moves.appendPawnMove(b, pos, pos+7)
			}
		}

		if !isRightBorder {
			// Right take
			sq := b.squares[pos+9]
			if sq != nil && sq.color == White {
				moves.appendPawnMove(b, pos, pos+9)
			}
		}

		if row == 4 {
			// En passant
			if !isRightBorder {
				rp := b.squares[pos+1]

				if rp != nil && rp.color == White && rp.kind == Pawn &&
					rp.flags&HasMovedRightBefore != 0 {
					moves.appendPawnMove(b, pos, pos+9)
				}
			}
			if !isLeftBorder {
				rp := b.squares[pos-1]

				if rp != nil && rp.color == White && rp.kind == Pawn &&
					rp.flags&HasMovedRightBefore != 0 {
					moves.appendPawnMove(b, pos, pos+7)
				}
			}
		}
	}

	return moves
}

func (b *Board) getKnightMoves(pos Position) []Move {
	var moves Moves

	top := pos.getSquaresFromTop()
	bottom := pos.getSquaresFromBottom()
	left := pos.getSquaresFromLeft()
	right := pos.getSquaresFromRight()

	if top >= 1 && left >= 2 {
		moves.Append(b, pos, pos-8-2)
	}
	if bottom >= 1 && left >= 2 {
		moves.Append(b, pos, pos+8-2)
	}
	if bottom >= 2 && left >= 1 {
		moves.Append(b, pos, pos+16-1)
	}
	if bottom >= 2 && right >= 1 {
		moves.Append(b, pos, pos+16+1)
	}
	if bottom >= 1 && right >= 2 {
		moves.Append(b, pos, pos+8+2)
	}
	if top >= 1 && right >= 2 {
		moves.Append(b, pos, pos-8+2)
	}
	if top >= 2 && right >= 1 {
		moves.Append(b, pos, pos-16+1)
	}
	if top >= 2 && left >= 1 {
		moves.Append(b, pos, pos-16-1)
	}

	return moves
}

func (b *Board) appendNextMove(from Position, to Position, moves *Moves) bool {
	piece := b.squares[to]

	if piece == nil {
		moves.Append(b, from, to)
		return true
	}

	if piece.color != b.turn {
		moves.Append(b, from, to)
	}

	return false
}

func (b *Board) getRookMoves(pos Position) []Move {
	var moves Moves

	for p := int(pos) - 8; p >= 0; p -= 8 {
		if b.appendNextMove(pos, Position(p), &moves) == false {
			break
		}
	}

	for p := int(pos) + 8; p <= 63; p += 8 {
		if b.appendNextMove(pos, Position(p), &moves) == false {
			break
		}
	}

	for p := pos - 1; p.getSquaresFromRight() != 0; p-- {
		if b.appendNextMove(pos, p, &moves) == false {
			break
		}
	}

	for p := pos + 1; p.getSquaresFromLeft() != 0; p++ {
		if b.appendNextMove(pos, p, &moves) == false {
			break
		}
	}

	return moves
}

func (b *Board) getBishopMoves(pos Position) []Move {
	var moves Moves

	if !pos.isTopBorder() && !pos.isLeftBorder() {
		for i := int(pos) - 9; i >= 0; i -= 9 {
			p := Position(i)

			if b.appendNextMove(pos, Position(i), &moves) == false {
				break
			}

			if p.isTopBorder() || p.isLeftBorder() {
				break
			}
		}
	}

	if !pos.isTopBorder() && !pos.isRightBorder() {
		for i := int(pos) - 7; i >= 0; i -= 7 {
			p := Position(i)

			if b.appendNextMove(pos, Position(i), &moves) == false {
				break
			}

			if p.isTopBorder() || p.isRightBorder() {
				break
			}
		}
	}

	if !pos.isBottomBorder() && !pos.isLeftBorder() {
		for i := int(pos) + 7; i <= 63; i += 7 {
			p := Position(i)

			if b.appendNextMove(pos, Position(i), &moves) == false {
				break
			}

			if p.isBottomBorder() || p.isLeftBorder() {
				break
			}
		}
	}

	if !pos.isBottomBorder() && !pos.isRightBorder() {
		for i := int(pos) + 9; i <= 63; i += 9 {
			p := Position(i)

			if b.appendNextMove(pos, Position(i), &moves) == false {
				break
			}

			if p.isBottomBorder() || p.isRightBorder() {
				break
			}
		}
	}

	return moves
}

func (b *Board) getQueenMoves(pos Position) []Move {
	return append(b.getBishopMoves(pos), b.getRookMoves(pos)...)
}

func (b *Board) getKingMoves(pos Position) []Move {
	var moves Moves

	if !pos.isLeftBorder() {
		moves.Append(b, pos, pos-1)
	}
	if !pos.isLeftBorder() && !pos.isTopBorder() {
		moves.Append(b, pos, pos-9)
	}
	if !pos.isTopBorder() {
		moves.Append(b, pos, pos-8)
	}
	if !pos.isTopBorder() && !pos.isRightBorder() {
		moves.Append(b, pos, pos-7)
	}
	if !pos.isRightBorder() {
		moves.Append(b, pos, pos+1)
	}
	if !pos.isBottomBorder() && !pos.isRightBorder() {
		moves.Append(b, pos, pos+9)
	}
	if !pos.isBottomBorder() {
		moves.Append(b, pos, pos+8)
	}
	if !pos.isBottomBorder() && !pos.isLeftBorder() {
		moves.Append(b, pos, pos+7)
	}

	if (pos == 4 || pos == 60) &&
		b.squares[pos].flags&HasMoved == 0 {
		// O-O
		if b.squares[pos+1] == nil &&
			b.squares[pos+2] == nil &&
			b.squares[pos+3] != nil &&
			b.squares[pos+3].flags&HasMoved == 0 &&
			!b.areSquaresAttaqued([]Position{pos, pos + 1, pos + 2}) {
			moves.Append(b, pos, pos+2)
		}

		// O-O-O
		if b.squares[pos-1] == nil &&
			b.squares[pos-2] == nil &&
			b.squares[pos-3] == nil &&
			b.squares[pos-4] != nil &&
			b.squares[pos-4].flags&HasMoved == 0 &&
			!b.areSquaresAttaqued([]Position{pos, pos - 1, pos - 2}) {
			moves.Append(b, pos, pos-2)
		}
	}

	return moves
}

func (b *Board) getMovesOpts(legal bool, kingMoves bool, cache bool) Moves {
	if cache && b.movesCache != nil {
		return b.movesCache
	}

	var moves Moves

	for i, square := range b.squares {
		if square == nil {
			continue
		}

		if square.color != b.turn {
			continue
		}

		pos := Position(i)

		switch square.kind {
		case Pawn:
			moves = append(moves, b.getPawnMoves(pos)...)

		case Knight:
			moves = append(moves, b.getKnightMoves(pos)...)

		case Bishop:
			moves = append(moves, b.getBishopMoves(pos)...)

		case Rook:
			moves = append(moves, b.getRookMoves(pos)...)

		case Queen:
			moves = append(moves, b.getQueenMoves(pos)...)

		case King:
			if kingMoves == true {
				moves = append(moves, b.getKingMoves(pos)...)
			}

		default:
			log.Panicf("unknown type: %d", square.kind)
		}
	}

	if !legal {
		return moves
	}

	var legalMoves Moves

	for _, move := range moves {
		if b.isMoveLegal(&move) {
			legalMoves = append(legalMoves, move)
		}
	}

	if cache == true {
		b.movesCache = legalMoves
	}

	return legalMoves
}

func (b *Board) GetMoves() Moves {
	return b.getMovesOpts(true, true, true)
}

func (b *Board) Move(move *Move) State {
	b.moveNoCheck(move)

	h := b.hash()
	val, ok := b.history[h]
	if ok {
		if val == 2 {
			return StateDrawByRepetition
		}

		b.history[h] = b.history[h] + 1
	} else {
		b.history[h] = 1
	}

	if b.hasEnoughMaterial() == false {
		return StateDrawByInsufficientMaterial
	}

	m := b.GetMoves()
	if len(m) > 0 {
		return StatePlaying
	}

	if !b.isCheck() {
		return StateDrawByStalemate
	}

	if b.turn == White {
		return StateBlackWins
	}

	return StateWhiteWins
}

func (b *Board) moveNoCheck(move *Move) {
	b.movesCache = nil

	isTake := b.squares[move.to] != nil

	b.squares[move.to] = b.squares[move.from]
	b.squares[move.from] = nil

	piece := b.squares[move.to]

	if move.promoteTo != Empty {
		piece.kind = move.promoteTo
	}

	for _, square := range b.squares {
		if square == nil {
			continue
		}

		piece.flags &= ^HasMovedRightBefore
	}
	piece.flags |= HasMoved | HasMovedRightBefore

	if piece.kind == King {
		if move.to-move.from == 2 {
			// O-O
			b.squares[move.to-1] = b.squares[move.to+1]
			b.squares[move.to+1] = nil
		} else if move.from-move.to == 2 {
			// O-O-O
			b.squares[move.to+1] = b.squares[move.to-2]
			b.squares[move.to-2] = nil
		}

		if piece.color == White {
			b.whiteKingCache = move.to
		} else {
			b.blackKingCache = move.to
		}
	} else if piece.kind == Pawn {
		// En passant
		if (move.to-move.from)%8 != 0 && isTake == false {
			if piece.color == White {
				b.squares[move.to+8] = nil
			} else {
				b.squares[move.to-8] = nil
			}
		}
	}

	b.turn.swap()
	b.nbMoves++
}
