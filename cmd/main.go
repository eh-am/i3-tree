package cmd

import (
	"context"
	"log"
	"os"
)

func Main() {
	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
