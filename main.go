// filestats counts files by extension for a given directory tree.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"filestats/internal/stats"
)

func main() {
	byLang := flag.Bool("l", false, "group results by language instead of extension")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: filestats [-l] [path]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	start := time.Now()

	result, err := stats.Analyze(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *byLang {
		result = stats.GroupByLanguage(result)
	}

	stats.Print(result)
	fmt.Printf("\nCompleted in %s\n", time.Since(start).Round(time.Millisecond))
}
