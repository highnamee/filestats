package stats

import (
	"fmt"
	"strings"
)

const (
	colFile = 9 // wide enough for "1,000,000"
	colSize = 10
	colPct  = 7
)

// Print writes a formatted table of file statistics to stdout.
// When r.GroupedByLanguage is true, Language is shown first and Extension(s) second.
func Print(r *Result) {
	if r.TotalFiles == 0 {
		fmt.Println("No files found.")
		return
	}

	extW, langW := computeWidths(r)

	if r.GroupedByLanguage {
		printGroupedByLanguage(r, extW, langW)
	} else {
		printGroupedByExt(r, extW, langW)
	}
}

// computeWidths returns the minimum column widths needed to fit all data without truncation.
func computeWidths(r *Result) (extW, langW int) {
	if r.GroupedByLanguage {
		extW = len("Extension(s)")
	} else {
		extW = len("Extension")
	}
	langW = len("Language")
	for _, stat := range r.Stats {
		if len(stat.Ext) > extW {
			extW = len(stat.Ext)
		}
		if len(stat.Language) > langW {
			langW = len(stat.Language)
		}
	}
	return
}

func printGroupedByExt(r *Result, extW, langW int) {
	fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n", extW, "Extension", langW, "Language", colFile, "Files", colSize, "Size", colPct, "Share")
	printSeparator(extW, langW)
	for _, stat := range r.Stats {
		fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n",
			extW, stat.Ext,
			langW, stat.Language,
			colFile, formatInt(stat.Files),
			colSize, formatBytes(stat.Bytes),
			colPct, formatPct(stat.Files, r.TotalFiles),
		)
	}
	printSeparator(extW, langW)
	fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n", extW, "Total", langW, "", colFile, formatInt(r.TotalFiles), colSize, formatBytes(r.TotalBytes), colPct, "100.0%")
}

func printGroupedByLanguage(r *Result, extW, langW int) {
	fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n", langW, "Language", extW, "Extension(s)", colFile, "Files", colSize, "Size", colPct, "Share")
	printSeparator(langW, extW)
	for _, stat := range r.Stats {
		fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n",
			langW, stat.Language,
			extW, stat.Ext,
			colFile, formatInt(stat.Files),
			colSize, formatBytes(stat.Bytes),
			colPct, formatPct(stat.Files, r.TotalFiles),
		)
	}
	printSeparator(langW, extW)
	fmt.Printf("%-*s  %-*s  %*s  %*s  %*s\n", langW, "Total", extW, "", colFile, formatInt(r.TotalFiles), colSize, formatBytes(r.TotalBytes), colPct, "100.0%")
}

func printSeparator(w1, w2 int) {
	fmt.Printf("%s  %s  %s  %s  %s\n",
		strings.Repeat("─", w1),
		strings.Repeat("─", w2),
		strings.Repeat("─", colFile),
		strings.Repeat("─", colSize),
		strings.Repeat("─", colPct),
	)
}

func formatInt(n int) string {
	s := fmt.Sprintf("%d", n)
	out := make([]byte, 0, len(s)+(len(s)-1)/3)
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, byte(c))
	}
	return string(out)
}

func formatPct(files, total int) string {
	return fmt.Sprintf("%.1f%%", float64(files)/float64(total)*100)
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
