package internal

import (
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/prune"
)

type PruneStratName string

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
		return nil, BadStratError{strat}
	}
}
