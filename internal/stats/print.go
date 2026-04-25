package stats

import "fmt"

const (
	colExt  = -20
	colLang = -16
	colFile = 8
	colSize = 10
	colPct  = 8
)

// Print writes a formatted table of file statistics to stdout.
// When r.GroupedByLanguage is true, Language is shown first and Extension(s) second.
func Print(r *Result) {
	if r.TotalFiles == 0 {
		fmt.Println("No files found.")
		return
	}

	if r.GroupedByLanguage {
		printGroupedByLanguage(r)
	} else {
		printGroupedByExt(r)
	}
}

func printGroupedByExt(r *Result) {
	fmt.Printf("%*s %*s %*s %*s %*s\n", colExt, "Extension", colLang, "Language", colFile, "Files", colSize, "Size", colPct, "Share")
	printSeparator()
	for _, stat := range r.Stats {
		pct := float64(stat.Files) / float64(r.TotalFiles) * 100
		fmt.Printf("%*s %*s %*d %*s %*.1f%%\n",
			colExt, stat.Ext,
			colLang, stat.Language,
			colFile, stat.Files,
			colSize, formatBytes(stat.Bytes),
			colPct-1, pct,
		)
	}
	printSeparator()
	fmt.Printf("%*s %*s %*d %*s %*.1f%%\n", colExt, "Total", colLang, "", colFile, r.TotalFiles, colSize, formatBytes(r.TotalBytes), colPct-1, 100.0)
}

func printGroupedByLanguage(r *Result) {
	fmt.Printf("%*s %*s %*s %*s %*s\n", colLang, "Language", colExt, "Extension(s)", colFile, "Files", colSize, "Size", colPct, "Share")
	printSeparator()
	for _, stat := range r.Stats {
		pct := float64(stat.Files) / float64(r.TotalFiles) * 100
		fmt.Printf("%*s %*s %*d %*s %*.1f%%\n",
			colLang, stat.Language,
			colExt, stat.Ext,
			colFile, stat.Files,
			colSize, formatBytes(stat.Bytes),
			colPct-1, pct,
		)
	}
	printSeparator()
	fmt.Printf("%*s %*s %*d %*s %*.1f%%\n", colLang, "Total", colExt, "", colFile, r.TotalFiles, colSize, formatBytes(r.TotalBytes), colPct-1, 100.0)
}

func printSeparator() {
	fmt.Printf("%*s %*s %*s %*s %*s\n",
		colExt, "────────────────────",
		colLang, "────────────────",
		colFile, "────────",
		colSize, "──────────",
		colPct, "───────",
	)
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
