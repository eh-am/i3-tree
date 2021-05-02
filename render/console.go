package render

import (
	"fmt"
	"io"

	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

type console struct {
	w  io.Writer
	au aurora.Aurora
}

func NewConsole(w io.Writer, colors bool) *console {
	return &console{
		w:  w,
		au: aurora.NewAurora(colors),
	}
}

func (t *console) Render(tree *i3.Tree) {
	t.print(tree.Root, "", "", 0)
}

func (t *console) print(node *i3.Node, prefix string, marker string, level int) {
	ftype := t.formatType(node, t.au)
	flayout := t.formatLayout(node, t.au)

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
