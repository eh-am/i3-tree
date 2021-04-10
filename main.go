package main

import (
	"log"
	"os"

	cr "github.com/eh-am/i3-tree-viewer/conrenderer"
	"github.com/eh-am/i3-tree-viewer/i3treeviewer"
	"github.com/eh-am/i3-tree-viewer/treepruner"
	"go.i3wm.org/i3/v4"
)

func main() {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	i3tv := i3treeviewer.NewI3TreeViewer(
		&treepruner.NeWsPruner{},
		cr.NewConRenderer(os.Stdout),
	)

	i3tv.View(&tree)
}
