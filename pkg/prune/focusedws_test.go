package prune_test

import (
	"testing"

	"github.com/eh-am/i3-tree/pkg/prune"
	"go.i3wm.org/i3/v4"
	"gotest.tools/assert"
)

func TestFocusedWorkspace(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		tree := &i3.Tree{}
		want := &i3.Tree{}

		w := &prune.FocusedWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

	t.Run("tree with a non empty focused workspace, ", func(t *testing.T) {
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
							{Type: "con", Name: "1", Focused: true},
							{Type: "con", Name: "2"},
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
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1", Focused: true},
							{Type: "con", Name: "2"},
						},
					},
				},
			},
		}

		w := &prune.NonEmptyWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

	t.Run("tree with multiple non empty workspaces and a focused one", func(t *testing.T) {
		tree := &i3.Tree{
			Root: &i3.Node{
				Type: "Root",
				Nodes: []*i3.Node{
					{
						Type:  "output",
						Nodes: []*i3.Node{{Type: "con", Name: "1"}},
					},
					{
						Name: "3",
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2", Focused: true},
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
							{Type: "con", Name: "2", Focused: true},
						},
					},
				},
			},
		}

		w := &prune.FocusedWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})
}
