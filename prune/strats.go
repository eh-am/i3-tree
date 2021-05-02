package prune

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

type PruneStratName string

var (
	NonEmptyWsStrat PruneStratName = "non-empty-ws"
	NoneStrat       PruneStratName = "none"

	AvailableStrats = []PruneStratName{
		NonEmptyWsStrat,
		NoneStrat,
	}
)

type Pruner interface {
	Prune(*i3.Tree) *i3.Tree
}

func NewStrat(s string) (pruner Pruner, err error) {
	switch PruneStratName(s) {
	case NoneStrat:
		pruner = &NoOpPruner{}

	case NonEmptyWsStrat:
		pruner = &NonEmptyWs{}

	default:
		// TODO err
		err = fmt.Errorf("invalid strat " + s)
	}

	return
}

func (p PruneStratName) String() string {
	return string(p)
}
