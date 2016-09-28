package geneticchess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sort"
	"time"
)

const DefaultFilePath = "/tmp/genetic-chess-phenotype.json"
const idLen = 8

type AI struct {
	Generation uint
	CloneID    string

	Genes map[string]float64

	tables *Tables `json:"-"`
}

type boardInfoPiecesCount struct {
	pawn   uint
	knight uint
	bishop uint
	rook   uint
	queen  uint

	total uint
}

type boardInfo struct {
	nbPieces int

	whiteCount boardInfoPiecesCount
	blackCount boardInfoPiecesCount
}

func NewAI() *AI {
	ai := &AI{
		Genes: map[string]float64{
			"PieceValuePawn":      1.0,
			"PieceValueKnight":    3.0,
			"PieceValueBishop":    3.1,
			"PieceValueRook":      5.0,
			"PieceValueQueen":     9.9,
			"PiecePositionPawn":   1.0,
			"PiecePositionKnight": 0.3,
			"PiecePositionBishop": 0.3,
			"PiecePositionRook":   0.5,
			"PiecePositionQueen":  1.0,
			"PiecePositionKing":   0.1,
			"NbMovesFactor":       0.01,
			"PruneRatio":          0.5,
			"MinKeptNodes":        5.0,
			"EndgameNbPieces":     10.0,
		},
		tables: NewTables(),
	}
	ai.CloneID = ai.getRandID(idLen)

	return ai
}

func NewAIEmpty() *AI {
	ai := &AI{
		tables: NewTables(),
		Genes:  map[string]float64{},
	}
	ai.CloneID = ai.getRandID(idLen)

	for _, gene := range genes {
		ai.Genes[gene.name] = gene.min
	}

	return ai
}

func NewAIRandom() *AI {
	ai := &AI{
		tables: NewTables(),
		Genes:  map[string]float64{},
	}
	ai.CloneID = ai.getRandID(idLen)

	for _, gene := range genes {
		ai.Genes[gene.name] = (gene.max - gene.min) * rand.Float64()
	}

	return ai
}

func NewAIFromJSON(data []byte) (*AI, error) {
	AI := &AI{}
	err := json.Unmarshal([]byte(data), AI)
	if err != nil {
		return nil, fmt.Errorf("invalid ai file: %v", err)
	}

	AI.tables = NewTables()

	return AI, nil
}

func NewAIFromFile(file string) (*AI, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	ai, err := NewAIFromJSON([]byte(data))
	if err != nil {
		return nil, err
	}

	return ai, nil
}

func (ai *AI) String() string {
	return fmt.Sprintf("G%d-%s", ai.Generation, ai.CloneID)
}

func (ai *AI) getGene(gene string) float64 {
	val, ok := ai.Genes[gene]
	if !ok {
		panic("cannot find gene " + gene)
	}

	return val
}

func (ai *AI) clone() *AI {
	clone := &AI{
		Genes:      make(map[string]float64),
		tables:     ai.tables,
		Generation: ai.Generation,
	}

	for key, val := range ai.Genes {
		clone.Genes[key] = val
	}

	return clone
}

type Job struct {
	node  *Node
	board *Board
}

type PrunableNode struct {
	node  *Node
	board *Board

	move Move
}

type PrunableNodes []PrunableNode

func (nm PrunableNodes) Len() int {
	return len(nm)
}

func (nm PrunableNodes) Less(i int, j int) bool {
	return nm[i].node.score <= nm[j].node.score
}

func (nm PrunableNodes) Swap(i int, j int) {
	nm[i], nm[j] = nm[j], nm[i]
}

func (ai *AI) buildNode(b *Board, move *Move, depth int) (*Node, State) {
	state := b.Move(move)

	switch state {
	case StateBlackWins:
		return &Node{score: -100000.0 + float64(depth), turn: b.turn}, state

	case StateWhiteWins:
		return &Node{score: 100000.0 - float64(depth), turn: b.turn}, state

	case StateDrawByRepetition:
		fallthrough
	case StateDrawByStalemate:
		fallthrough
	case StateDrawByInsufficientMaterial:
		return &Node{score: 0, turn: b.turn}, state

	case StatePlaying:
		return &Node{
			turn:     b.turn,
			score:    ai.evalPosition(b),
			children: make(map[Move]*Node),
		}, state
	}

	panic("unknown state: " + string(state))
	return nil, 0
}

func (ai *AI) pruneNodes(list PrunableNodes, turn Color,
	pruneRatio float64, minKeptNodes uint) PrunableNodes {
	if minKeptNodes == 0 {
		// Keep at least one node.
		minKeptNodes = 1
	}

	sort.Sort(list)

	to := len(list)
	from := int(float64(to) * pruneRatio)
	if to-from < int(minKeptNodes) {
		from = to - int(minKeptNodes)
		if from < 0 {
			from = 0
		}
	}
	toKeep := to - from

	if turn == White {
		return list[from:]
	}
	return list[:toKeep]
}

func (ai *AI) buildTree(b *Board, timeToThink time.Duration,
	maxAllowedDepth uint, truncate bool) *Node {
	end := time.Now().Add(timeToThink)
	root := &Node{children: map[Move]*Node{}, turn: b.turn}
	todo := []*Job{&Job{node: root, board: b}}
	pruneRatio := ai.getGene("PruneRatio")
	minKeptNodes := math.Floor(ai.getGene("MinKeptNodes") + 0.5)
	maxDepth := 0

loop:
	for len(todo) > 0 {
		job := todo[0]
		todo = todo[1:]

		var list PrunableNodes

		depth := job.board.nbMoves - b.nbMoves
		m := job.board.GetMoves()
		for _, move := range m {
			newBoard := job.board.clone()

			if depth > maxDepth {
				maxDepth = depth
				if maxAllowedDepth > 0 &&
					maxDepth > int(maxAllowedDepth) {
					break loop
				}
			}

			child, state := ai.buildNode(newBoard, &move, depth)
			if state == StatePlaying {
				list = append(list, PrunableNode{
					move:  move,
					node:  child,
					board: newBoard,
				})
			} else {
				job.node.children[move] = child
			}
		}

		if len(list) > 0 {
			list = ai.pruneNodes(list, job.board.turn,
				pruneRatio, uint(minKeptNodes))
			if len(list) == 0 {
				panic("all moves were pruned")
			}

			for _, elem := range list {
				job.node.children[elem.move] = elem.node
				todo = append(todo, &Job{
					node:  elem.node,
					board: elem.board,
				})
			}
		}

		if timeToThink > 0 && time.Now().After(end) {
			break loop
		}
	}

	if truncate == true && maxDepth > 1 {
		// Truncate level that could not be completely computed.
		root.truncate(maxDepth - 1)
	}

	return root
}

func (ai *AI) GetBestMoveScore(b *Board,
	timeToThink time.Duration, maxDepth uint) (*Move, float64) {
	tree := ai.buildTree(b, timeToThink, maxDepth, false)
	tree.factorize()

	return tree.best, tree.score
}

func (ai *AI) GetBestMove(b *Board, timeToThink time.Duration,
	maxDepth uint) *Move {
	move, _ := ai.GetBestMoveScore(b, timeToThink, maxDepth)
	return move
}

func (ai *AI) evalPieces(b *Board, info *boardInfo) float64 {
	score := 0.0

	for _, piece := range b.squares {
		if piece == nil || piece.kind == King {
			continue
		}

		val := ai.getGene("PieceValue" + piece.kind.GetName())
		score += val * piece.color.score()

		info.nbPieces++

		var counter *boardInfoPiecesCount

		if piece.color == White {
			counter = &info.whiteCount
		} else {
			counter = &info.blackCount
		}

		switch piece.kind {
		case Pawn:
			counter.pawn++
		case Knight:
			counter.knight++
		case Bishop:
			counter.bishop++
		case Rook:
			counter.rook++
		case Queen:
			counter.queen++
		}
	}

	info.whiteCount.total =
		info.whiteCount.pawn +
			info.whiteCount.knight +
			info.whiteCount.bishop +
			info.whiteCount.rook +
			info.whiteCount.queen
	info.blackCount.total =
		info.blackCount.pawn +
			info.blackCount.knight +
			info.blackCount.bishop +
			info.blackCount.rook +
			info.blackCount.queen

	return score
}

func (ai *AI) evalPositionnbMoves(b *Board, info *boardInfo) float64 {
	var whiteMoves int
	var blackMoves int

	if b.turn == White {
		whiteMoves = len(b.GetMoves())

		b.turn.swap()
		blackMoves = len(b.getMovesOpts(true, true, false))
		b.turn.swap()
	} else {
		blackMoves = len(b.GetMoves())

		b.turn.swap()
		whiteMoves = len(b.getMovesOpts(true, true, false))
		b.turn.swap()
	}

	// Ignore winning-side number of moves.
	if info.whiteCount.total == 0 {
		blackMoves = 0
	} else if info.blackCount.total == 0 {
		whiteMoves = 0
	}

	return float64(whiteMoves-blackMoves) * ai.getGene("NbMovesFactor")
}

func (ai *AI) evalPiecesPosition(b *Board,
	ignoreColor *Color, beginnng bool) float64 {
	var tables *TablesStage
	var score float64

	if beginnng == true {
		tables = ai.tables.Beginning
	} else {
		tables = ai.tables.Endgame
	}

	pawnFactor := ai.getGene("PiecePositionPawn")
	knightFactor := ai.getGene("PiecePositionKnight")
	bishopFactor := ai.getGene("PiecePositionBishop")
	rookFactor := ai.getGene("PiecePositionRook")
	queenFactor := ai.getGene("PiecePositionQueen")
	kingFactor := ai.getGene("PiecePositionKing")

	for pos, piece := range b.squares {
		if piece == nil {
			continue
		}

		if ignoreColor != nil &&
			*ignoreColor == piece.color &&
			piece.kind != Pawn {
			// Make sure to never ignore pawns
			continue
		}

		var tbc *TablesColor
		var res float64

		if piece.color == White {
			tbc = tables.White
		} else {
			tbc = tables.Black
		}

		switch piece.kind {
		case Pawn:
			res = tbc.Pawn[pos] * pawnFactor
		case Knight:
			res = tbc.Knight[pos] * knightFactor
		case Bishop:
			res = tbc.Bishop[pos] * bishopFactor
		case Rook:
			res = tbc.Rook[pos] * rookFactor
		case Queen:
			res = tbc.Queen[pos] * queenFactor
		case King:
			res = tbc.King[pos] * kingFactor
		}

		if piece.color == White {
			score += res
		} else {
			score -= res
		}
	}

	return score
}

func (ai *AI) evalPosition(b *Board) float64 {
	var info boardInfo
	var score float64

	score += ai.evalPieces(b, &info)

	if b.nbMoves > 30 {
		score += ai.evalPositionnbMoves(b, &info)
	}

	if info.whiteCount.total == 0 {
		// Ignore black piece position to get a faster checkmate
		color := Black
		score += ai.evalPiecesPosition(b, &color, false)
	} else if info.blackCount.total == 0 {
		// Ignore white piece position to get a faster checkmate
		color := White
		score += ai.evalPiecesPosition(b, &color, false)
	} else {
		endgame := uint(math.Floor(ai.getGene("EndgameNbPieces") + 0.5))
		score += ai.evalPiecesPosition(b, nil,
			info.whiteCount.total+info.blackCount.total > endgame)
	}

	return score
}

func (ai *AI) getResult(p1 *AI, p2 *AI, score float64) [2]*Result {
	res := [2]*Result{
		&Result{Player: p1},
		&Result{Player: p2},
	}

	switch score {
	case -1:
		res[1].score = 1
	case 0:
		res[1].score = 0.5
		res[0].score = 0.5
	case 1:
		res[0].score = 1
	}

	return res
}

func (ai *AI) Play(opponent *AI, color Color,
	timeToThink time.Duration, maxDepth uint, results chan [2]*Result) {
	b := NewBoard()
	factor := color.score()

	for {
		var state State
		var move *Move

		if b.nbMoves > 200 {
			// No need to waste more time.
			results <- ai.getResult(ai, opponent, 0.0)
			break
		}

		if b.turn == color {
			move = ai.GetBestMove(b, timeToThink, maxDepth)
		} else {
			move = opponent.GetBestMove(b, timeToThink, maxDepth)
		}

		state = b.Move(move)
		if state == StatePlaying {
			continue
		}

		switch state {
		case StateDrawByRepetition:
			fallthrough
		case StateDrawByStalemate:
			fallthrough
		case StateDrawByInsufficientMaterial:
			results <- ai.getResult(ai, opponent, 0.0)
		case StateWhiteWins:
			results <- ai.getResult(ai, opponent, 1.0*factor)
		case StateBlackWins:
			results <- ai.getResult(ai, opponent, -1.0*factor)
		default:
			panic(fmt.Sprintf("unknown state %v", state))
		}

		break
	}
}

func (ai *AI) getRandID(n uint) string {
	chars := []rune(
		"abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"0123456789")

	id := make([]rune, n)
	for i := uint(0); i < n; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}

	return string(id)
}

func (ai *AI) mute(nbGenes uint, size float64) *AI {
	clone := ai.clone()
	clone.Generation++

	if len(clone.CloneID) < 4 {
		clone.CloneID = ai.getRandID(4)
	}

	clone.CloneID = ai.getRandID(4) + ai.CloneID[0:4]

	selectedGenes := make(map[int]bool)

	for uint(len(selectedGenes)) < nbGenes {
		selectedGenes[rand.Int()%len(genes)] = true
	}

	for i := range selectedGenes {
		gene := genes[i]
		val := ai.getGene(gene.name)

		for {
			diff := (0.5 - rand.Float64()) * (gene.max - gene.min)

			newVal := val + diff
			newVal = math.Min(newVal, gene.max)
			newVal = math.Max(newVal, gene.min)

			if newVal == val {
				continue
			}

			ai.Genes[gene.name] = newVal
			break
		}
	}

	return clone
}
