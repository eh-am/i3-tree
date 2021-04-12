package conrenderer

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

func formatLayout(layout i3.Layout, au aurora.Aurora) string {
	s := string(layout)
	switch layout {
	case "stacked":
		return au.BrightGreen(s).String()
	case "tabbed":
		return au.Green(s).String()
	case "splith":
		return au.BrightYellow(s).String()
	case "splitv":
		return au.Yellow(s).String()
	default:
		return s
	}
}

func formatType(t i3.NodeType, au aurora.Aurora) string {
	s := string(t)

	switch t {
	case "workspace":
		return au.Cyan(s).String()
	case "con":
		return au.Blue(s).String()
	case "output":
		return au.Magenta(s).String()
	default:
		return s
	}
}

func wrapBrackets(s string) string {
	if s != "" {
		return fmt.Sprint("[", s, "]")
	}
	return s
}

func generateLayout(node *i3.Node, au aurora.Aurora) string {
	s := ""
	// only show layout if it has children
	if len(node.Nodes) > 0 {
		return wrapBrackets(formatLayout(node.Layout, au))
	}

	return s
}

func generateType(node *i3.Node, au aurora.Aurora) string {
	return wrapBrackets(formatType(node.Type, au))
}
