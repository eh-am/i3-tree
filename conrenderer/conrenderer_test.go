package conrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/eh-am/i3-tree-viewer/conrenderer"
	"github.com/eh-am/i3-tree-viewer/testutils"
	"github.com/hexops/autogold"
)

func Test_conRenderer_Render(t *testing.T) {
	type args struct {
		filepath string
		color    bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"should render correctly with colors",
			args{
				filepath: "../testdata/pruned_tree.json",
				color:    true,
			},
		},
		{
			"should render correctly without colors",
			args{
				filepath: "../testdata/pruned_tree.json",
				color:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := testutils.TreeFromJSONFile(tt.args.filepath)

			var writer bytes.Buffer
			r := conrenderer.NewConRenderer(io.Writer(&writer), tt.args.color)
			r.Render(tree)
			//
			got := writer.String()
			autogold.Equal(t, got)
		})
	}
}
