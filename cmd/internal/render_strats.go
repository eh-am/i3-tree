package internal

import (
	"os"

	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/render"
)

type RendererStrat string

var (
	// Console strategy
	ConsoleStrat RendererStrat = "console"
	// Console, but no color strategy
	ConsoleNoColorStrat RendererStrat = "no-color"

	// List of all available render strategies
	AvailableRendererStrats = []RendererStrat{
		ConsoleStrat,
		ConsoleNoColorStrat,
	}
)

// NewRenderer creates a i3treeviewer.Renderer
// Based on a strategy
// Otherwise it fails with BadStratError
func NewRenderer(strat string) (i3treeviewer.Renderer, error) {
	switch RendererStrat(strat) {
	case ConsoleStrat:
		return render.NewColoredConsole(os.Stdout), nil

	case ConsoleNoColorStrat:
		return render.NewMonochromaticConsole(os.Stdout), nil

	default:
		return nil, BadStratError{strat}
	}
}
