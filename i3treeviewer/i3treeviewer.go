package i3treeviewer

import (
	"go.i3wm.org/i3/v4"
)

type Pruner interface {
	Prune(*i3.Tree) *i3.Tree
}
type Renderer interface {
	Render(*i3.Tree)
}

type i3TreeViewer struct {
	Pruner
	Renderer
}

func NewI3TreeViewer(Pruner Pruner, Renderer Renderer) i3TreeViewer {
	return i3TreeViewer{
		Pruner,
		Renderer,
	}
}

func (i3tv *i3TreeViewer) View(tree *i3.Tree) {
	n := i3tv.Prune(tree)

	i3tv.Render(n)
}
