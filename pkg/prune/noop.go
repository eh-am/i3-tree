package prune

import "go.i3wm.org/i3/v4"

type NoOp struct{}

func (w *NoOp) Prune(tree *i3.Tree) *i3.Tree {
	return tree
}
