package prune

import (
	"go.i3wm.org/i3/v4"
)

type FocusedWs struct{}

func (w *FocusedWs) Prune(tree *i3.Tree) *i3.Tree {
	if tree.Root != nil {
		// found a focused workspace
		i, _ := w.walk(tree.Root)

		if i != "" {
			a := Ws{WsIndex: i}
			return a.Prune(tree)
		}
	}

	// there's nothing focused so don't show anything
	return &i3.Tree{
		Root: nil,
	}
}

// return the name of the workspace with a focused window
func (w *FocusedWs) walk(src *i3.Node) (string, bool) {
	if src == nil {
		return "", false
	}

	if src.Focused {
		return "", true
	}

	// if any of the children are focused
	for _, n := range src.Nodes {
		id, focused := w.walk(n)
		if id != "" {
			return id, true
		}

		if focused {
			if src.Type == "workspace" {
				return src.Name, true
			}
			return "", true
		}
	}

	return "", false
}
