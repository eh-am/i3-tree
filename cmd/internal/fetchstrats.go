package internal

import (
	"github.com/eh-am/i3-tree-viewer/pkg/fetch"
	"github.com/eh-am/i3-tree-viewer/pkg/i3treeviewer"
)

type BadFetchStratError struct{ StratName string }

func (e BadFetchStratError) Error() string {
	return "invalid fetch strat: " + e.StratName
}

// Fetch Strategies
type FetchStratName string

var (
	FromI3 FetchStratName = "i3"
)

func NewFetcher(strat string) (i3treeviewer.Fetcher, error) {
	switch FetchStratName(strat) {
	case FromI3:
		return fetch.FromI3{}, nil
	default:
		return nil, BadFetchStratError{strat}
	}
}