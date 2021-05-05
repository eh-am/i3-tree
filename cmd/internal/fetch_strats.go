package internal

import (
	"github.com/eh-am/i3-tree/pkg/fetch"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
)

// Fetch Strategies
type FetchStratName string

var (
	FromI3 FetchStratName = "i3"
	Mock   FetchStratName = "mock"

	AvailableFetchStrats = []FetchStratName{
		FromI3,
		Mock,
	}
)

func NewFetcher(strat string) (i3treeviewer.Fetcher, error) {
	switch FetchStratName(strat) {
	case FromI3:
		return fetch.FromI3{}, nil
	case Mock:
		return fetch.FromFake{}, nil
	default:
		return nil, BadStratError{strat}
	}
}
