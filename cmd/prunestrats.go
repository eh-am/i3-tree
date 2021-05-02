package cmd

import (
	"github.com/eh-am/i3-tree-viewer/i3treeviewer"
	"github.com/eh-am/i3-tree-viewer/prune"
)

type PruneStratName string

type BadPruneStratError struct {
	StratName string
}

func (e BadPruneStratError) Error() string {
	return "invalid prune strat: " + e.StratName
}

var (
	NonEmptyWsPruneStrat PruneStratName = "non-empty-ws"
	NonePruneStrat       PruneStratName = "none"

	AvailablePruneStrats = []PruneStratName{
		NonEmptyWsPruneStrat,
		NonePruneStrat,
	}
)

// NewPruner decides which prune strategy to use
// Based on the flag name
func NewPruner(strat string) (i3treeviewer.Pruner, error) {
	switch PruneStratName(strat) {
	case NonePruneStrat:
		return &prune.NoOp{}, nil

	case NonEmptyWsPruneStrat:
		return &prune.NonEmptyWs{}, nil

	default:
		return nil, BadPruneStratError{strat}
	}
}
