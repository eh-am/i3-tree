package treepruner

import "go.i3wm.org/i3/v4"

type NoOpPruner struct{}

func (w *NoOpPruner) Prune(tree *i3.Tree) *i3.Tree {
	return tree
}
