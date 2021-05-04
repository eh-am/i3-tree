package internal_test

import (
	"testing"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/render"
	"github.com/stretchr/testify/assert"
)

func TestNewRenderer(t *testing.T) {
	cases := []struct {
		stratName string
		want      i3treeviewer.Renderer
		wantErr   error
	}{
		{"console", render.ColoredConsole{}, nil},
		{"no-color", render.MonochromaticConsole{}, nil},
		{"unknown", nil, internal.BadStratError{"unknown"}},
	}

	for _, tt := range cases {
		t.Run(tt.stratName, func(t *testing.T) {
			got, gotErr := internal.NewRenderer(tt.stratName)

			assert.IsType(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
