package geneticchess

type Position uint8

var borders = map[Position]bool{
	0:  true,
	1:  true,
	2:  true,
	3:  true,
	4:  true,
	5:  true,
	6:  true,
	8:  true,
	7:  true,
	15: true,
	16: true,
	23: true,
	24: true,
	31: true,
	32: true,
	39: true,
	40: true,
	47: true,
	48: true,
	55: true,
	56: true,
	57: true,
	58: true,
	59: true,
	60: true,
	61: true,
	62: true,
	63: true,
}

const (
	InitialPositionBlackRookA   Position = 0
	InitialPositionBlackKnightB Position = 1
	InitialPositionBlackBishopC Position = 2
	InitialPositionBlackQueen   Position = 3
	InitialPositionBlackBishopF Position = 5
	InitialPositionBlackKnightG Position = 6
	InitialPositionBlackRookH   Position = 7

	InitialPositionWhiteRookA   Position = 56
	InitialPositionWhiteKnightB Position = 57
	InitialPositionWhiteBishopC Position = 58
	InitialPositionWhiteQueen   Position = 59
	InitialPositionWhiteBishopF Position = 61
	InitialPositionWhiteKnightG Position = 62
	InitialPositionWhiteRookH   Position = 63
)

func (p Position) getRow() int {
	return int(p) / 8
}

func (p Position) getCol() int {
	return int(p) % 8
}

func (p Position) getSquaresFromTop() int {
	return p.getRow()
}

func (p Position) getSquaresFromBottom() int {
	return 7 - p.getRow()
}

func (p Position) getSquaresFromLeft() int {
	return int(p) % 8
}

func (p Position) getSquaresFromRight() int {
	return 7 - int(p)%8
}

func (p Position) isTopBorder() bool {
	return p <= 7
}

func (p Position) isBottomBorder() bool {
	return p >= 56
}

func (p Position) isLeftBorder() bool {
	return p%8 == 0
}

func (p Position) isRightBorder() bool {
	return p%8 == 7
}
