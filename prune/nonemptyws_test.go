package prune_test

import (
	"testing"

	"github.com/eh-am/i3-tree-viewer/prune"
	"go.i3wm.org/i3/v4"
	"gotest.tools/assert"
)

func TestNonEmptyWorkspace(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		tree := &i3.Tree{}
		want := &i3.Tree{}

		w := &prune.NonEmptyWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

	t.Run("tree with empty workspace", func(t *testing.T) {
		tree := &i3.Tree{
			Root: &i3.Node{
				Type:  "workspace",
				Nodes: []*i3.Node{},
			},
		}
		want := &i3.Tree{}

		w := &prune.NonEmptyWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

	t.Run("tree with a non empty workspace", func(t *testing.T) {
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

		want := &i3.Tree{
			Root: &i3.Node{
				Type: "Root",
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
		}

		w := &prune.NonEmptyWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})

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
					// this should still exist
					{
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},

					// this should still exist
					{
						Type: "workspace",
						Nodes: []*i3.Node{
							{Type: "con", Name: "1"},
							{Type: "con", Name: "2"},
						},
					},

					// should still exist
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

		w := &prune.NonEmptyWs{}
		got := w.Prune(tree)

		assert.DeepEqual(t, want, got)
	})
}
