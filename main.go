package main

import (
	"fmt"
	"log"
	"os"
	//	"strings"

	"github.com/logrusorgru/aurora/v3"
	"go.i3wm.org/i3/v4"
)

func main() {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	n := prune(tree.Root)

	printTree(n, "", "")
}

// prune a subtree and generate a tree
// where workspace is not empty
func prune(src *i3.Node) *i3.Node {
	if src.Type == "workspace" && len(src.Nodes) > 0 {
		return src
	}

	subtrees := make([]*i3.Node, 0)
	for _, n := range src.Nodes {
		r := prune(n)
		if r != nil {
			subtrees = append(subtrees, r)
		}
	}

	if len(subtrees) > 0 {
		src.Nodes = subtrees
		return src
	}

	return nil
}

func colorizeLayout(layout i3.Layout) aurora.Value {
	s := string(layout)
	switch layout {
	case "stacked":
		return aurora.BrightGreen(s)
	case "tabbed":
		return aurora.Green(s)
	case "splith":
		return aurora.BrightYellow(s)
	case "splitv":
		return aurora.Yellow(s)
	default:
		// TODO noop
		return aurora.BgBrightBlack(s)
	}
}

func colorizeType(t i3.NodeType) aurora.Value {
	s := string(t)

	switch t {
	case "workspace":
		return aurora.Cyan(s)
	case "con":
		return aurora.Blue(s)
	case "root":
		return aurora.White(s)
	case "output":
		return aurora.Magenta(s)
	default:
		// TODO noop
		return aurora.White(s)
	}
}

// first thing is to normalize this tree
func printTree(node *i3.Node, prefix string, marker string) {
	// TODO
	// filter values i don't want
	coloredType := colorizeType(node.Type)

	coloredLayout := ""

	// only show layout if it's
	if len(node.Nodes) > 0 {
		coloredLayout = fmt.Sprint("[", colorizeLayout(node.Layout), "]")
	}

	// TODO remove me
	if string(node.Type) == string(node.Layout) {
		coloredLayout = ""
	}

	output := fmt.Sprintf("%s%s[%s]%s %s", prefix, marker, coloredType, coloredLayout, node.Name)
	fmt.Fprint(os.Stdout, output, "\n")

	// TODO

	for i, n := range node.Nodes {
		newPrefix := prefix + ""
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
			newPrefix = newPrefix + "   "
		}

		printTree(n, newPrefix, newMarker)
	}
}
