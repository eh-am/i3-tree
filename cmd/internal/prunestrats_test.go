package internal_test

import (
	"testing"

	"github.com/eh-am/i3-tree-viewer/cmd/internal"
	"github.com/eh-am/i3-tree-viewer/prune"
	"github.com/stretchr/testify/assert"
)

func TestNewPruner(t *testing.T) {
	t.Run("strat: none", func(t *testing.T) {
		want := &prune.NoOp{}
		got, err := internal.NewPruner("none")

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("strat: non-empty-ws", func(t *testing.T) {
		want := &prune.NonEmptyWs{}
		got, err := internal.NewPruner("non-empty-ws")

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("strat: badstrat", func(t *testing.T) {
		_, err := internal.NewPruner("badstrat")

		wantErr := internal.BadPruneStratError{"badstrat"}
		assert.Equal(t, wantErr, err)

		// Not entirely sure about testing the error message
		// But since it will be shown to clients
		// I will consider part of the API and therefore should tested
		assert.Equal(t, wantErr.Error(), "invalid prune strat: badstrat")
	})
}
