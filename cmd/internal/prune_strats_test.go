package internal_test

import (
	"testing"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/prune"
	"github.com/stretchr/testify/assert"
)

func TestNewPruner(t *testing.T) {
	cases := []struct {
		stratName string
		want      i3treeviewer.Pruner
		wantErr   error
	}{
		{"none", &prune.NoOp{}, nil},
		{"non-empty-ws", &prune.NonEmptyWs{}, nil},
		{"unknown", nil, internal.BadStratError{"unknown"}},
	}

	for _, tt := range cases {
		t.Run(tt.stratName, func(t *testing.T) {
			got, gotErr := internal.NewPruner(tt.stratName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
