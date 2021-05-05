package render

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"go.i3wm.org/i3/v4"
)

func (t *console) wrapBrackets(s string) string {
	if s != "" {
		return fmt.Sprint("[", s, "]")
	}
	return s
}

func (t *console) formatLayout(node *i3.Node, au aurora.Aurora) string {
	if node == nil {
		return ""
	}

	formatFn := func(layout i3.Layout, au aurora.Aurora) string {
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

	s := ""

	// only show layout if it has children
	if len(node.Nodes) > 0 {
		return t.wrapBrackets(formatFn(node.Layout, au))
	}

	return s
}

func (t *console) formatType(node *i3.Node, au aurora.Aurora) string {
	if node == nil {
		return ""
	}

	formatFn := func(t i3.NodeType, au aurora.Aurora) string {
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

	return t.wrapBrackets(formatFn(node.Type, au))
}
