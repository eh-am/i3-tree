package prune_test

import (
	"testing"

	"github.com/eh-am/i3-tree-viewer/pkg/prune"
	"go.i3wm.org/i3/v4"
	"gotest.tools/assert"
)

func TestNoop(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		tree := &i3.Tree{}
		want := &i3.Tree{}

		w := &prune.NoOp{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

	t.Run("slightly more complex tree", func(t *testing.T) {
		tree := &i3.Tree{
			Root: &i3.Node{
				Type: "Root",
				Nodes: []*i3.Node{
					// this should be pruned
					{
						Type:  "output",
						Nodes: []*i3.Node{{Type: "con", Name: "1"}},
					},

					// this should still exist
					{
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},
				},
			},
		}

		want := tree

		w := &prune.NoOp{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})
}
