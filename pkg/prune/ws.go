package prune

import "go.i3wm.org/i3/v4"

// Ws prunes the tree maintaining a specific workspace
type Ws struct {
	WsIndex string
}

func (w *Ws) Prune(tree *i3.Tree) *i3.Tree {
	return pruneSubtree(tree, func(src *i3.Node) bool {
		// there's a workspace
		// and there are things within it
		if src != nil && src.Type == "workspace" && src.Name == w.WsIndex {
			return true
		}
		return false
	})
}
