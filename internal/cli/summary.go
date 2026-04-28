package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// RunSummary holds effective CLI options for the configuration footer.
type RunSummary struct {
	Root    string
	ByLang  bool
	JSONOut bool
	Top     int
	OutFile string
	Exclude []string
	LOC     bool
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

func formatDefaultExcludes(root string) string {
	if _, err := os.Stat(filepath.Join(root, ".gitignore")); err != nil {
		return ".git/"
	}
	return ".git/; Files/Folders in root .gitignore"
}

// PrintRunSummary writes the "--- Configuration summary" block used after a run.
func PrintRunSummary(w io.Writer, s RunSummary) error {
	ew := &errWriter{w: w}
	ew.printf("\n---\nConfiguration summary:\n")
	ew.printf("  Path:              %s\n", s.Root)
	if s.ByLang {
		ew.printf("  Group by:          language (-l)\n")
	} else {
		ew.printf("  Group by:          extension\n")
	}
	if s.JSONOut {
		ew.printf("  Output format:     JSON (stdout, -json)\n")
	} else {
		ew.printf("  Output format:     table\n")
	}
	if s.Top > 0 {
		ew.printf("  Limit:             top %d (-top)\n", s.Top)
	} else {
		ew.printf("  Limit:             all rows (-top=0)\n")
	}
	if s.OutFile != "" {
		abs, err := filepath.Abs(s.OutFile)
		if err != nil {
			abs = s.OutFile
		}
		ew.printf("  JSON file:         %s (-o)\n", abs)
	} else {
		ew.printf("  JSON file:         (none)\n")
	}
	ew.printf("  Default excludes:  %s\n", formatDefaultExcludes(s.Root))
	if len(s.Exclude) > 0 {
		ew.printf("  Extra excludes:    %s (-exclude)\n", strings.Join(s.Exclude, ", "))
	}
	if s.LOC {
		ew.printf("  Lines of code:     counted (-loc=true)\n")
	} else {
		ew.printf("  Lines of code:     skipped (-loc=false)\n")
	}
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
