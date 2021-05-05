package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var flagHelp = `i3-tree generates a user friendly view of the i3 tree

EXAMPLES
# display focused workspace
i3-tree

# display all non empty workspaces
i3-tree all

# show a specific workspace (for example, workspace 6)
i3-tree 6

# show focused workspace, with no colors
i3-tree --render=no-color

# use mock data (useful if you don't have i3 running)
i3-tree --from=mock
`

var fetchStratName *string
var renderStratName *string

var rootFs *flag.FlagSet
var root *ffcli.Command

func init() {
	rootFs = flag.NewFlagSet("root", flag.ExitOnError)

	fetchStratName = rootFs.String(
		"from",
		string(internal.FromI3),
		"where to fetch the tree from. available: "+fmt.Sprintf("%s", internal.AvailableFetchStrats),
	)

	renderStratName = rootFs.String(
		"render",
		string(internal.ConsoleStrat), // Default
		"where/how to render the output to. available: "+fmt.Sprintf("%s", internal.AvailableRendererStrats),
	)

	root = &ffcli.Command{
		Name:       "i3-tree",
		ShortUsage: "i3-tree",
		LongHelp:   flagHelp,
		ShortHelp:  "Print the i3 tree in a user friendly format",
		FlagSet:    rootFs,
		Exec:       rootExec,
	}
}

func rootExec(ctx context.Context, args []string) error {
	fetcher, err := internal.NewFetcher(*fetchStratName)
	if err != nil {
		return err
	}

	pruneArg := ""
	if len(args) > 0 {
		pruneArg = args[0]
	}
	pruner, err := internal.NewPruner(pruneArg)
	if err != nil {
		return err
	}

	renderer, err := internal.NewRenderer(*renderStratName)
	if err != nil {
		return err
	}

	i3tv := i3treeviewer.NewI3TreeViewer(
		fetcher,
		pruner,
		renderer,
	)

	return i3tv.View()
}
