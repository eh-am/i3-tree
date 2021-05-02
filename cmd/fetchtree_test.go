package cmd_test

import (
	"testing"

	"github.com/eh-am/i3-tree-viewer/cmd"
	"github.com/eh-am/i3-tree-viewer/fetch"
	"github.com/eh-am/i3-tree-viewer/i3treeviewer"
	"github.com/stretchr/testify/assert"
)

func TestNewFetcher(t *testing.T) {
	cases := []struct {
		stratName string
		want      i3treeviewer.Fetcher
		wantErr   error
	}{
		{"i3", fetch.FromI3{}, nil},
		{"unknown", nil, cmd.BadFetchStratError{"unknown"}},
	}

	for _, tt := range cases {
		t.Run(tt.stratName, func(t *testing.T) {
			got, gotErr := cmd.NewFetcher(tt.stratName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
