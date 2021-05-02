package i3treeviewer

import (
	"go.i3wm.org/i3/v4"
)

type Fetcher interface {
	Fetch() (i3.Tree, error)
}
type Pruner interface {
	Prune(*i3.Tree) *i3.Tree
}
type Renderer interface {
	Render(*i3.Tree)
}

type i3TreeViewer struct {
	Fetcher
	Pruner
	Renderer
}

func NewI3TreeViewer(Fetcher Fetcher, Pruner Pruner, Renderer Renderer) i3TreeViewer {
	return i3TreeViewer{
		Fetcher,
		Pruner,
		Renderer,
	}
}

func (i3tv *i3TreeViewer) View() error {
	tree, err := i3tv.Fetcher.Fetch()
	if err != nil {
		// TODO
		return err
	}
	n := i3tv.Prune(&tree)

	i3tv.Render(n)
	return nil
}
