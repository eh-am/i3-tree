package conrenderer

import (
	"fmt"
	"io"

	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

type conRenderer struct {
	w  io.Writer
	au aurora.Aurora
}

func NewConRenderer(w io.Writer, colors bool) *conRenderer {
	return &conRenderer{
		w:  w,
		au: aurora.NewAurora(colors),
	}
}

func (t *conRenderer) Render(tree *i3.Tree) {
	t.print(tree.Root, "", "", 0)
}

func (t *conRenderer) print(node *i3.Node, prefix string, marker string, level int) {
	ftype := generateType(node, t.au)
	flayout := generateLayout(node, t.au)

	fmt.Fprint(
		t.w,
		prefix,
		marker,
		ftype,
		flayout,
		" ",
		node.Name,
		"\n",
	)

	for i, n := range node.Nodes {
		newPrefix := prefix
		newMarker := ""

		// figure out what's the marker for the next iteration
		if i == len(node.Nodes)-1 {
			newMarker = "└──" // last node
		} else {
			newMarker = "├──" // middle node
		}

		// i am currently a middle node
		if marker == "├──" {
			// so my children should display my trunk
			newPrefix = newPrefix + "│  "
		} else {
			// don't ident starting from root
			if level == 0 {
				newPrefix = ""
			} else {
				newPrefix = newPrefix + "   "
			}
		}

		t.print(n, newPrefix, newMarker, level+1)
	}
}
