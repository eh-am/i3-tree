package conrenderer

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

func formatLayout(layout i3.Layout) string {
	s := string(layout)
	switch layout {
	case "stacked":
		return aurora.BrightGreen(s).String()
	case "tabbed":
		return aurora.Green(s).String()
	case "splith":
		return aurora.BrightYellow(s).String()
	case "splitv":
		return aurora.Yellow(s).String()
	case "output":
		return ""
	default:
		return s
	}
}

func formatType(t i3.NodeType) string {
	s := string(t)

	switch t {
	case "workspace":
		return aurora.Cyan(s).String()
	case "con":
		return aurora.Blue(s).String()
	case "output":
		return aurora.Magenta(s).String()
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

func generateLayout(node *i3.Node) string {
	s := ""
	// only show layout if it has children
	if len(node.Nodes) > 0 {
		return wrapBrackets(formatLayout(node.Layout))
	}

	return s
}

func generateType(node *i3.Node) string {
	return wrapBrackets(formatType(node.Type))
}
