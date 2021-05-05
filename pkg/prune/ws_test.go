package prune_test

import (
	"testing"

	"github.com/eh-am/i3-tree/pkg/prune"
	"go.i3wm.org/i3/v4"
	"gotest.tools/assert"
)

func TestWorkspace(t *testing.T) {
	t.Run("tree with multiple non empty workspaces", func(t *testing.T) {
		tree := &i3.Tree{
			Root: &i3.Node{
				Type: "Root",
				Nodes: []*i3.Node{
					// this should be pruned
					{
						Type:  "output",
						Nodes: []*i3.Node{{Type: "con", Name: "1"}},
					},
					{
						Name: "3",
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},
					{
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},
					{
						Type: "output",
						Nodes: []*i3.Node{
							{
								Type: "workspace",
								Nodes: []*i3.Node{
									{Type: "con", Name: "1"},
									{Type: "con", Name: "2"},
								},
							},
						},
					},
				},
			},
		}

		want := &i3.Tree{
			Root: &i3.Node{
				Type: "Root",
				Nodes: []*i3.Node{
					{
						Name: "3",
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},
				},
			},
		}

		w := &prune.Ws{"3"}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})
}
