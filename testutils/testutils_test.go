package testutils

import (
	"github.com/hexops/autogold"
	"testing"
)

func TestTreeFromJSONFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"should load correctly",
			args{
				filepath: "../testdata/tree1.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TreeFromJSONFile(tt.args.filepath)
			autogold.Equal(t, got)
		})
	}
}
