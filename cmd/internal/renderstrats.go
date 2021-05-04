package internal

import (
	"os"

	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/eh-am/i3-tree/pkg/render"
)

type RendererStrat string

type BadRendererStratError struct {
	StratName string
}

func (e BadRendererStratError) Error() string {
	return "invalid render strat: " + e.StratName
}

var (
	ConsoleStrat        RendererStrat = "console"
	ConsoleNoColorStrat RendererStrat = "console-no-color"

	AvailableRendererStrats = []RendererStrat{
		ConsoleStrat,
		ConsoleNoColorStrat,
	}
)

func NewRenderer(strat string) (i3treeviewer.Renderer, error) {
	switch RendererStrat(strat) {
	case ConsoleStrat:
		return render.NewColoredConsole(os.Stdout), nil

	case ConsoleNoColorStrat:
		return render.NewMonochromaticConsole(os.Stdout), nil

	default:
		return nil, BadRendererStratError{strat}
	}
}
