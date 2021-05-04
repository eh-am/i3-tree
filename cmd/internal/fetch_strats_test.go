package internal_test

import (
	"testing"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/fetch"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/stretchr/testify/assert"
)

func TestNewFetcher(t *testing.T) {
	cases := []struct {
		stratName string
		want      i3treeviewer.Fetcher
		wantErr   error
	}{
		{"i3", fetch.FromI3{}, nil},
		{"fake", fetch.FromFake{}, nil},
		{"unknown", nil, internal.BadStratError{"unknown"}},
	}

	for _, tt := range cases {
		t.Run(tt.stratName, func(t *testing.T) {
			got, gotErr := internal.NewFetcher(tt.stratName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
