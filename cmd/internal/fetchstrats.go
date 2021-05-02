package internal

import (
	"fmt"

	"github.com/eh-am/i3-tree-viewer/pkg/fetch"
	"github.com/eh-am/i3-tree-viewer/pkg/i3treeviewer"
)

type BadFetchStratError struct{ StratName string }

func (e BadFetchStratError) Error() string {
	return "invalid fetch strat: " + e.StratName + "\n" +
		"possible values: " + fmt.Sprintf("%s", AvailableFetchStrats)
}

// Fetch Strategies
type FetchStratName string

var (
	FromI3 FetchStratName = "i3"
	Fake   FetchStratName = "fake"

	AvailableFetchStrats = []FetchStratName{
		FromI3,
		Fake,
	}
)

func NewFetcher(strat string) (i3treeviewer.Fetcher, error) {
	switch FetchStratName(strat) {
	case FromI3:
		return fetch.FromI3{}, nil
	case Fake:
		return fetch.FromFake{}, nil
	default:
		return nil, BadFetchStratError{strat}
	}
}
