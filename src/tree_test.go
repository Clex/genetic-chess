package geneticchess

import (
	"testing"
)

func TestTreeTruncateTree(t *testing.T) {
	getMaxDepth := func(tree *Node) int {
		max := 0

		tree.WalkNodes(func(node *Node, depth int) {
			if depth > max {
				max = depth
			}
		}, 0)

		// "max" is the deepest node, therefore the deepest leaf
		// is one level deeper.
		return max + 1
	}

	tree := &Node{
		turn: White,
		children: map[Move]*Node{
			Move{from: 1, to: 2}: &Node{
				turn:  Black,
				score: 1,
			},
			Move{from: 1, to: 3}: &Node{
				turn: Black,
				children: map[Move]*Node{
					Move{from: 2, to: 1}: &Node{
						turn:  White,
						score: -3,
					},
					Move{from: 2, to: 2}: &Node{
						turn:  White,
						score: -4,
						children: map[Move]*Node{
							Move{from: 42, to: 1}: &Node{
								turn:  Black,
								score: 10000,
							},
						},
					},
				},
			},
			Move{from: 1, to: 4}: &Node{
				turn:  Black,
				score: 3,
			},
			Move{from: 1, to: 5}: &Node{
				turn: Black,
				children: map[Move]*Node{
					Move{from: 2, to: 3}: &Node{
						turn:  White,
						score: 42,
					},
					Move{from: 2, to: 4}: &Node{
						turn:  White,
						score: -42,
					},
				},
			},
		},
	}

	tree.truncate(2)
	if getMaxDepth(tree) != 2 {
		t.Fatalf("expected depth of %d instead of %d",
			2, getMaxDepth(tree))
	}
}

func TestTreeFactorize(t *testing.T) {
	tree := &Node{
		turn: White,
		children: map[Move]*Node{
			Move{from: 1, to: 2}: &Node{
				turn:  Black,
				score: 1,
			},
			Move{from: 1, to: 3}: &Node{
				turn: Black,
				children: map[Move]*Node{
					Move{from: 2, to: 1}: &Node{
						turn:  White,
						score: -3,
					},
					Move{from: 2, to: 2}: &Node{
						turn:  White,
						score: -4,
					},
				},
			},
			Move{from: 1, to: 4}: &Node{
				turn:  Black,
				score: 3,
			},
			Move{from: 1, to: 5}: &Node{
				turn: Black,
				children: map[Move]*Node{
					Move{from: 2, to: 3}: &Node{
						turn:  White,
						score: 42,
					},
					Move{from: 2, to: 4}: &Node{
						turn:  White,
						score: -42,
					},
				},
			},
		},
	}

	tree.factorize()

	if tree.best.from != 1 || tree.best.to != 4 {
		t.Fatalf("expected best move to be 1:4 instead of %s", tree.best.String())
	}

	if tree.score != 3 {
		t.Fatalf("expected score to be 3 instead of %d", tree.score)
	}
}
