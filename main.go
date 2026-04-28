// filestats counts files by extension for a given directory tree.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"filestats/internal/cli"
	"filestats/internal/stats"
)

// version is set at build time via: -ldflags "-X main.version=x.y.z".
var version = "dev"

func main() {
	byLang := flag.Bool("l", false, "group results by language instead of extension")
	jsonOut := flag.Bool("json", false, "print results as JSON to stdout instead of table")
	outFile := flag.String("o", "", "save results as JSON to `file`")
	showVersion := flag.Bool("version", false, "print version and exit")
	top := flag.Int("top", 0, "show only the top `N` results (0 = all)")
	var excludes cli.StringsFlag
	flag.Var(&excludes, "exclude", "exclude files/dirs matching `pattern` (glob; repeatable, comma-separated)")
	loc := flag.Bool("loc", true, "count lines of code (disable with -loc=false for faster runs)")
	gitignore := flag.Bool("respect-gitignore", true, "respect root .gitignore (disable with -respect-gitignore=false to include ignored files)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "filestats — count files by extension or language in a directory tree.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n  filestats [flags] [path]\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  filestats .                           count by extension in current dir\n")
		fmt.Fprintf(os.Stderr, "  filestats -l /src                     group by language\n")
		fmt.Fprintf(os.Stderr, "  filestats -top 10 -json -o out.json   top 10 results as JSON, also save to file\n")
		fmt.Fprintf(os.Stderr, "  filestats -respect-gitignore=false .  include files normally ignored by .gitignore\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		if err := cli.PrintFlagDefaults(os.Stderr); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}

	os.Args = cli.ReorderArgs(os.Args, map[string]bool{"l": true, "json": true, "version": true, "loc": true, "respect-gitignore": true})
	flag.Parse()

	if *showVersion {
		fmt.Printf("filestats %s\n", version)
		return
	}

	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	start := time.Now()

	excludePatterns := []string(excludes)
	result, err := stats.Analyze(root, excludePatterns, *loc, *gitignore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *byLang {
		result = stats.GroupByLanguage(result)
	}

	result = stats.TopN(result, *top)

	summary := cli.RunSummary{
		Root:      root,
		ByLang:    *byLang,
		JSONOut:   *jsonOut,
		Top:       *top,
		OutFile:   *outFile,
		Exclude:   excludePatterns,
		LOC:       *loc,
		Gitignore: *gitignore,
	}

	var summaryOut io.Writer = os.Stdout
	if *jsonOut {
		if err := stats.WriteJSON(os.Stdout, result); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		summaryOut = os.Stderr
	} else {
		stats.Print(result)
		fmt.Println()
	}

	if *outFile != "" {
		if err := stats.SaveJSON(*outFile, result); err != nil {
			fmt.Fprintf(os.Stderr, "error writing %s: %v\n", *outFile, err)
			os.Exit(1)
		}
	}

	if err := cli.PrintRunFooter(summaryOut, summary, time.Since(start)); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
