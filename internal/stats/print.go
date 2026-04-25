package stats

import (
	"fmt"
	"strings"
)

const (
	colFile = 9
	colSize = 10
	colPct  = 7
)

// Print writes a formatted, coloured table of file statistics to stdout.
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

// computeWidths returns column widths for ext and lang.
// When grouped by language, extW is capped so the full row fits the terminal.
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
	if r.GroupedByLanguage && extW > maxExtWidth {
		extW = maxExtWidth
	}
	return
}

func printRow(cols ...string) {
	fmt.Println(strings.Join(cols, "  "))
}

func printGroupedByExt(r *Result, extW, langW int) {
	printRow(
		styleHeader.Sprint(padRight("Extension", len("Extension"), extW)),
		styleHeader.Sprint(padRight("Language", len("Language"), langW)),
		styleHeader.Sprint(padLeft("Files", len("Files"), colFile)),
		styleHeader.Sprint(padLeft("Size", len("Size"), colSize)),
		styleHeader.Sprint(padLeft("Share", len("Share"), colPct)),
	)
	printSeparator(extW, langW)
	for _, stat := range r.Stats {
		filesStr := formatInt(stat.Files)
		sizeStr := formatBytes(stat.Bytes)
		pct := float64(stat.Files) / float64(r.TotalFiles) * 100
		printRow(
			padRight(stat.Ext, len(stat.Ext), extW),
			padRight(coloredLang(stat.Language), len(stat.Language), langW),
			padLeft(filesStr, len(filesStr), colFile),
			padLeft(sizeStr, len(sizeStr), colSize),
			padLeft(coloredShare(pct), len(formatPct(stat.Files, r.TotalFiles)), colPct),
			percentBar(pct),
		)
	}
	printSeparator(extW, langW)
	totalFilesStr := formatInt(r.TotalFiles)
	totalSizeStr := formatBytes(r.TotalBytes)
	printRow(
		styleTotal.Sprint(padRight("Total", len("Total"), extW)),
		padRight("", 0, langW),
		styleTotal.Sprint(padLeft(totalFilesStr, len(totalFilesStr), colFile)),
		styleTotal.Sprint(padLeft(totalSizeStr, len(totalSizeStr), colSize)),
		styleTotal.Sprint(padLeft("100.0%", len("100.0%"), colPct)),
		percentBar(100),
	)
}

func printGroupedByLanguage(r *Result, extW, langW int) {
	printRow(
		styleHeader.Sprint(padRight("Language", len("Language"), langW)),
		styleHeader.Sprint(padRight("Extension(s)", len("Extension(s)"), extW)),
		styleHeader.Sprint(padLeft("Files", len("Files"), colFile)),
		styleHeader.Sprint(padLeft("Size", len("Size"), colSize)),
		styleHeader.Sprint(padLeft("Share", len("Share"), colPct)),
	)
	printSeparator(langW, extW)
	for _, stat := range r.Stats {
		filesStr := formatInt(stat.Files)
		sizeStr := formatBytes(stat.Bytes)
		pct := float64(stat.Files) / float64(r.TotalFiles) * 100
		extLines := wrapExts(stat.Ext, extW)
		printRow(
			padRight(coloredLang(stat.Language), len(stat.Language), langW),
			padRight(extLines[0], len(extLines[0]), extW),
			padLeft(filesStr, len(filesStr), colFile),
			padLeft(sizeStr, len(sizeStr), colSize),
			padLeft(coloredShare(pct), len(formatPct(stat.Files, r.TotalFiles)), colPct),
			percentBar(pct),
		)
		for _, line := range extLines[1:] {
			printRow(
				padRight("", 0, langW),
				padRight(styleDim.Sprint(line), len(line), extW),
				padLeft("", 0, colFile),
				padLeft("", 0, colSize),
				padLeft("", 0, colPct),
			)
		}
	}
	printSeparator(langW, extW)
	totalFilesStr := formatInt(r.TotalFiles)
	totalSizeStr := formatBytes(r.TotalBytes)
	printRow(
		styleTotal.Sprint(padRight("Total", len("Total"), langW)),
		padRight("", 0, extW),
		styleTotal.Sprint(padLeft(totalFilesStr, len(totalFilesStr), colFile)),
		styleTotal.Sprint(padLeft(totalSizeStr, len(totalSizeStr), colSize)),
		styleTotal.Sprint(padLeft("100.0%", len("100.0%"), colPct)),
		percentBar(100),
	)
}

func printSeparator(w1, w2 int) {
	printRow(
		styleDim.Sprint(strings.Repeat("─", w1)),
		styleDim.Sprint(strings.Repeat("─", w2)),
		styleDim.Sprint(strings.Repeat("─", colFile)),
		styleDim.Sprint(strings.Repeat("─", colSize)),
		styleDim.Sprint(strings.Repeat("─", colPct)),
		styleDim.Sprint(strings.Repeat("─", barWidth)),
	)
}

func formatPct(files, total int) string {
	return fmt.Sprintf("%.1f%%", float64(files)/float64(total)*100)
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
