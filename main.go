package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	cr "github.com/eh-am/i3-tree-viewer/conrenderer"
	"github.com/eh-am/i3-tree-viewer/i3treeviewer"
	"github.com/eh-am/i3-tree-viewer/treepruner"
	"github.com/peterbourgon/ff/v3/ffcli"
	"go.i3wm.org/i3/v4"
)

func main() {
	rootFs := flag.NewFlagSet("root", flag.ExitOnError)
	pruneStratName := rootFs.String("prune-strat", treepruner.NonEmptyWsStrat.String(), "what to remove from the raw tree i3 returns. available: "+fmt.Sprintf("%s", treepruner.AvailableStrats))

	root := &ffcli.Command{
		Name:       "i3-tree-viewer",
		ShortUsage: "i3-tree-viewer",
		ShortHelp:  "Print a tree of non-empty workspaces to console",
		FlagSet:    rootFs,
		Exec: func(ctx context.Context, args []string) error {
			tree, err := i3.GetTree()
			if err != nil {
				log.Fatal(err)
			}

			pruner, err := treepruner.NewStrat(*pruneStratName)
			if err != nil {
				return err
			}

			i3tv := i3treeviewer.NewI3TreeViewer(
				pruner,
				cr.NewConRenderer(os.Stdout, true),
			)

			i3tv.View(&tree)
			return nil
		},
	}

	root.ParseAndRun(context.Background(), os.Args[1:])
}
