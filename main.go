// filestats counts files by extension for a given directory tree.
package main

import (
	"fmt"
	"os"

	"filestats/internal/stats"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	result, err := stats.Analyze(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	stats.Print(result)
}
