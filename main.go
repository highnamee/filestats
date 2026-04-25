// filestats counts files by extension for a given directory tree.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"filestats/internal/cli"
	"filestats/internal/stats"
)

func main() {
	byLang := flag.Bool("l", false, "group results by language instead of extension")
	jsonOut := flag.Bool("json", false, "print results as JSON to stdout instead of table")
	outFile := flag.String("o", "", "save results as JSON to `file`")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: filestats [flags] [path]\n\n")
		flag.PrintDefaults()
	}

	os.Args = cli.ReorderArgs(os.Args, map[string]bool{"l": true, "json": true})
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

	if *jsonOut {
		if err := stats.WriteJSON(os.Stdout, result); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	stats.Print(result)

	if *outFile != "" {
		if err := stats.SaveJSON(*outFile, result); err != nil {
			fmt.Fprintf(os.Stderr, "error writing %s: %v\n", *outFile, err)
			os.Exit(1)
		}
		abs, err := filepath.Abs(*outFile)
		if err != nil {
			abs = *outFile
		}
		fmt.Printf("JSON saved to %s\n", abs)
	}

	fmt.Printf("Completed in %s\n", time.Since(start).Round(time.Millisecond))
}
