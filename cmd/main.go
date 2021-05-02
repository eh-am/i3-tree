package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eh-am/i3-tree-viewer/i3treeviewer"
	"github.com/eh-am/i3-tree-viewer/render"
	"github.com/peterbourgon/ff/v3/ffcli"
	"go.i3wm.org/i3/v4"
)

func Main() {
	rootFs := flag.NewFlagSet("root", flag.ExitOnError)
	pruneStratName :=
		rootFs.String(
			"prune-strat",
			string(NonEmptyWsPruneStrat), // Default
			"what to prune from the raw tree i3. "+fmt.Sprintf("%s", AvailablePruneStrats),
		)

	root := &ffcli.Command{
		Name:       "i3-tree-viewer",
		ShortUsage: "i3-tree-viewer",
		ShortHelp:  "Print the i3 tree in a user friendly format",
		FlagSet:    rootFs,
		Exec: func(ctx context.Context, args []string) error {
			tree, err := i3.GetTree()
			if err != nil {
				log.Fatal(err)
			}

			pruner, err := NewPruner(*pruneStratName)
			if err != nil {
				return err
			}

			i3tv := i3treeviewer.NewI3TreeViewer(
				pruner,
				render.NewConsole(os.Stdout, true),
			)

			i3tv.View(&tree)
			return nil
		},
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
