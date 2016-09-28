package geneticchess

import (
	"fmt"
	"math"
)

type Node struct {
	score    float64
	turn     Color
	best     *Move
	children map[Move]*Node
}

func (node *Node) WalkNodes(cb func(*Node, int), depth int) {
	if node.children == nil || len(node.children) == 0 {
		return
	}

	for _, child := range node.children {
		if child.children != nil && len(child.children) > 0 {
			for _, child := range node.children {
				child.WalkNodes(cb, depth+1)
			}
		}
	}

	cb(node, depth)
}

func (node *Node) WalkNodesBreatdth(cb func(*Node, int), depth int) {
	if node.children == nil || len(node.children) == 0 {
		return
	}

	cb(node, depth)

	// In case cb() removed children.
	if node.children == nil {
		return
	}

	for _, child := range node.children {
		if child.children != nil && len(child.children) > 0 {
			for _, child := range node.children {
				child.WalkNodes(cb, depth+1)
			}
		}
	}
}

func (node *Node) truncate(maxDepth int) {
	node.WalkNodesBreatdth(func(node *Node, depth int) {
		if depth >= maxDepth {
			node.children = nil
		}
	}, 0)
}

func (node *Node) factorize() {
	for {
		if node.children == nil || len(node.children) == 0 {
			break
		}

		node.WalkNodes(func(node *Node, depth int) {
			if node.turn == White {
				var maxMove *Move
				max := -10000000.0

				for move, child := range node.children {
					if child.score > max {
						max = child.score
						maxMove = &Move{
							from:      move.from,
							to:        move.to,
							promoteTo: move.promoteTo,
						}
					}
				}
				node.score = max
				node.best = maxMove
			} else {
				var minMove *Move
				min := math.MaxFloat64

				for move, child := range node.children {
					if child.score < min {
						min = child.score
						minMove = &Move{
							from:      move.from,
							to:        move.to,
							promoteTo: move.promoteTo,
						}
					}
				}
				node.score = min
				node.best = minMove
			}

			if node.best == nil {
				panic("cannot find best move")
			}

			node.children = nil
		}, 0)
	}
}

func (node *Node) Dump() {
	var recDump func(*Node, []string, int)

	recDump = func(node *Node, parents []string, depth int) {
		for move, child := range node.children {
			for i := 0; i < depth; i++ {
				fmt.Printf("=")
			}
			fmt.Printf(" ")
			for _, pMove := range parents {
				fmt.Printf("%s | ", pMove)
			}
			fmt.Printf("%s (%s) %f\n", move.String(), child.turn.String(), child.score)
			recDump(child, append(parents, move.String()), depth+1)
		}
	}

	recDump(node, []string{}, 1)
}
