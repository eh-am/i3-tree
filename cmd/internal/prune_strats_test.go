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
		arg     string
		want    i3treeviewer.Pruner
		wantErr error
	}{
		{"all", &prune.NonEmptyWs{}, nil},
		{"", &prune.FocusedWs{}, nil},
		{"raw", &prune.NoOp{}, nil},
		{"5", &prune.Ws{WsIndex: "5"}, nil},
	}

	for _, tt := range cases {
		t.Run(tt.arg, func(t *testing.T) {
			got, gotErr := internal.NewPruner(tt.arg)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
