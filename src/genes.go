package geneticchess

type Gene struct {
	name string
	min  float64
	max  float64
}

var genes = []Gene{
	/* Fitness function */
	Gene{name: "PieceValueQueen", min: 0.1, max: 10.0},
	Gene{name: "PieceValueRook", min: 0.1, max: 10.0},
	Gene{name: "PieceValueBishop", min: 0.1, max: 10.0},
	Gene{name: "PieceValueKnight", min: 0.1, max: 10.0},
	Gene{name: "PieceValuePawn", min: 0.1, max: 10.0},
	Gene{name: "PiecePositionQueen", min: 0.0, max: 1.0},
	Gene{name: "PiecePositionRook", min: 0.0, max: 1.0},
	Gene{name: "PiecePositionBishop", min: 0.0, max: 1.0},
	Gene{name: "PiecePositionKnight", min: 0.0, max: 1.0},
	Gene{name: "PiecePositionPawn", min: 0.0, max: 1.0},
	Gene{name: "PiecePositionKing", min: 0.0, max: 1.0},
	Gene{name: "NbMovesFactor", min: 0.0, max: 1.0},
	Gene{name: "EndgameNbPieces", min: 2.0, max: 16.0},

	/* Alpha-beta pruning strategy */
	Gene{name: "PruneRatio", min: 0.0, max: 0.99},
	Gene{name: "MinKeptNodes", min: 1.0, max: 5.0},
}
