package cli

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"
)

// RunSummary holds effective CLI options for the configuration footer.
type RunSummary struct {
	Root      string
	ByLang    bool
	JSONOut   bool
	Top       int
	OutFile   string
	Exclude   []string
	LOC       bool
	Gitignore bool
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) printf(format string, args ...any) {
	if ew.err != nil {
		return
	}
	_, ew.err = fmt.Fprintf(ew.w, format, args...)
}

func (ew *errWriter) row(key, value string) {
	ew.printf("  %-12s  %s\n", key, value)
}

// PrintRunSummary writes the configuration summary block after a run.
func PrintRunSummary(w io.Writer, s RunSummary) error {
	ew := &errWriter{w: w}

	ew.printf("\n--- Configuration ---\n")

	ew.row("Path", s.Root)

	if s.ByLang {
		ew.row("Group by", "language")
	} else {
		ew.row("Group by", "extension")
	}

	if s.JSONOut {
		ew.row("Output", "JSON")
	} else {
		ew.row("Output", "table")
	}

	if s.Top > 0 {
		ew.row("Limit", fmt.Sprintf("top %d", s.Top))
	} else {
		ew.row("Limit", "all rows")
	}

	if s.OutFile != "" {
		abs, err := filepath.Abs(s.OutFile)
		if err != nil {
			abs = s.OutFile
		}
		ew.row("JSON file", abs)
	} else {
		ew.row("JSON file", "(none)")
	}

	excludes := ".git/"
	if s.Gitignore {
		excludes += ", root .gitignore"
	}
	if len(s.Exclude) > 0 {
		excludes += ", " + strings.Join(s.Exclude, ", ")
	}
	ew.row("Excludes", excludes)

	if s.LOC {
		ew.row("Count LOC", "yes")
	} else {
		ew.row("Count LOC", "no")
	}

	ew.printf("---------------------\n")
	return ew.err
}

// PrintCompletedIn writes a timing line after the configuration summary.
func PrintCompletedIn(w io.Writer, elapsed time.Duration) error {
	_, err := fmt.Fprintf(w, "Completed in %s\n", elapsed.Round(time.Millisecond))
	return err
}

// PrintRunFooter writes the configuration summary and the completed-in line to w.
func PrintRunFooter(w io.Writer, summary RunSummary, elapsed time.Duration) error {
	if err := PrintRunSummary(w, summary); err != nil {
		return err
	}
	return PrintCompletedIn(w, elapsed)
}
