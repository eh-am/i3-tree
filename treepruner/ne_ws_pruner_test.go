package treepruner_test

import (
	"reflect"
	"testing"

	"github.com/eh-am/i3-tree-viewer/testutils"
	"github.com/eh-am/i3-tree-viewer/treepruner"
	"github.com/hexops/autogold"
	"go.i3wm.org/i3/v4"
)

func TestNeWsPruner_Prune(t *testing.T) {
	type args struct {
		tree *i3.Tree
	}
	tests := []struct {
		name string
		args args
		want *i3.Tree
	}{
		{
			"empty tree",
			args{
				tree: &i3.Tree{},
			},
			&i3.Tree{},
		},
		{
			"tree with only empty workspaces",
			args{
				tree: &i3.Tree{
					Root: &i3.Node{
						Type:  "workspace",
						Nodes: []*i3.Node{},
					},
				},
			},
			&i3.Tree{},
		},

		{
			"it should prune tree with a single non-empty workspace",
			args{
				tree: &i3.Tree{
					Root: &i3.Node{
						Type: "Root",
						Nodes: []*i3.Node{
							// this should be pruned
							&i3.Node{
								Type: "output",
								Nodes: []*i3.Node{
									&i3.Node{
										Type: "con",
										Name: "1",
									},
								},
							},

							// this should still exist
							&i3.Node{
								Type: "workspace",
								Nodes: []*i3.Node{
									&i3.Node{
										Type: "con",
										Name: "1",
									},
									&i3.Node{
										Type: "con",
										Name: "2",
									},
								},
							},
						},
					},
				},
			},
			&i3.Tree{
				Root: &i3.Node{
					Type: "Root",
					Nodes: []*i3.Node{
						&i3.Node{
							Type: "workspace",
							Nodes: []*i3.Node{
								&i3.Node{
									Type: "con",
									Name: "1",
								},
								&i3.Node{
									Type: "con",
									Name: "2",
								},
							},
						},
					},
				},
			},
		},

		{
			"it should prune tree with multiple non-empty workspaces",
			args{
				tree: &i3.Tree{
					Root: &i3.Node{
						Type: "Root",
						Nodes: []*i3.Node{
							// this should be pruned
							&i3.Node{
								Type: "output",
								Nodes: []*i3.Node{
									&i3.Node{Type: "con", Name: "1"},
								},
							},
							// this should still exist
							&i3.Node{
								Type: "workspace",
								Nodes: []*i3.Node{
									&i3.Node{Type: "con", Name: "1"},
									&i3.Node{Type: "con", Name: "2"},
								},
							},

							// this should still exist
							&i3.Node{
								Type: "workspace",
								Nodes: []*i3.Node{
									&i3.Node{Type: "con", Name: "1"},
									&i3.Node{Type: "con", Name: "2"},
								},
							},

							// should still exist
							&i3.Node{
								Type: "output",
								Nodes: []*i3.Node{
									&i3.Node{
										Type: "workspace",
										Nodes: []*i3.Node{
											&i3.Node{Type: "con", Name: "1"},
											&i3.Node{Type: "con", Name: "2"},
										},
									},
								},
							},
						},
					},
				},
			},
			&i3.Tree{
				Root: &i3.Node{
					Type: "Root",
					Nodes: []*i3.Node{
						// this should still exist
						&i3.Node{
							Type: "workspace",
							Nodes: []*i3.Node{
								&i3.Node{Type: "con", Name: "1"},
								&i3.Node{Type: "con", Name: "2"},
							},
						},

						// this should still exist
						&i3.Node{
							Type: "workspace",
							Nodes: []*i3.Node{
								&i3.Node{Type: "con", Name: "1"},
								&i3.Node{Type: "con", Name: "2"},
							},
						},

						&i3.Node{
							Type: "output",
							Nodes: []*i3.Node{
								&i3.Node{
									Type: "workspace",
									Nodes: []*i3.Node{
										&i3.Node{Type: "con", Name: "1"},
										&i3.Node{Type: "con", Name: "2"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &treepruner.NeWsPruner{}
			if got := w.Prune(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NeWsPruner.Prune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeWsPruner_PruneRealTree(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"real tree 1",
			args{
				filepath: "../testdata/tree1.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &treepruner.NeWsPruner{}
			tree := testutils.TreeFromJSONFile(tt.args.filepath)
			got := w.Prune(tree)
			autogold.Equal(t, got)
		})
	}
}
