package internal

import (
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/prune"
)

// NewPruner decides which prune strategy to use
// Based on the flag name
func NewPruner(arg string) (i3treeviewer.Pruner, error) {
	switch arg {
	case "":
		return &prune.FocusedWs{}, nil

	case "all":
		return &prune.NonEmptyWs{}, nil

	case "raw":
		return &prune.NoOp{}, nil

	default:
		return &prune.Ws{WsIndex: arg}, nil
	}
}
