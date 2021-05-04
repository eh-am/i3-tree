package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/eh-am/i3-tree/cmd/internal"
	"github.com/eh-am/i3-tree/pkg/i3treeviewer"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var pruneStratName *string
var fetchStratName *string
var renderStratName *string

var rootFs *flag.FlagSet
var root *ffcli.Command

func init() {
	rootFs = flag.NewFlagSet("root", flag.ExitOnError)

	fetchStratName = rootFs.String(
		"fetch-strat",
		string(internal.FromI3),
		"where to fetch the tree from. Available: "+fmt.Sprintf("%s", internal.AvailableFetchStrats),
	)

	pruneStratName = rootFs.String(
		"prune-strat",
		string(internal.NonEmptyWsPruneStrat), // Default
		"what to prune from the (possible raw) tree i3. Available: "+fmt.Sprintf("%s", internal.AvailablePruneStrats),
	)

	renderStratName = rootFs.String(
		"render-strat",
		string(internal.ConsoleStrat), // Default
		"where/how to render the output. Available: "+fmt.Sprintf("%s", internal.AvailableRendererStrats),
	)

	root = &ffcli.Command{
		Name:       "i3-tree",
		ShortUsage: "i3-tree",
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
