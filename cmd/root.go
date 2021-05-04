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
# show all non empty workspaces
i3-tree

# show all non empty workspaces, without colored output
i3-tree --render=no-color
`

var pruneStratName *string
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

	pruneStratName = rootFs.String(
		"pick",
		string(internal.NonEmptyWsPruneStrat), // Default
		fmt.Sprintf(`what to prune from the i3 tree. available: %s`, internal.AvailablePruneStrats),
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

	pruner, err := internal.NewPruner(*pruneStratName)
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
