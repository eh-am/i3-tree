package prune

import "go.i3wm.org/i3/v4"

// Non Empty Workspace Pruner
type NonEmptyWs struct{}

func (w *NonEmptyWs) Prune(tree *i3.Tree) *i3.Tree {
	return pruneSubtree(tree, func(src *i3.Node) bool {
		// there's a workspace
		// and there are things within it
		if src != nil && src.Type == "workspace" && len(src.Nodes) > 0 {
			return true
		}
		return false
	})
}
