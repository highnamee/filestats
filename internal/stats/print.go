package stats

import "fmt"

// Print writes a formatted table of file statistics to stdout.
func Print(r *Result) {
	if r.TotalFiles == 0 {
		fmt.Println("No files found.")
		return
	}

	fmt.Printf("%-20s %8s %10s %8s\n", "Extension", "Files", "Size", "Share")
	fmt.Printf("%-20s %8s %10s %8s\n", "─────────────────", "─────────", "──────────", "───────")

	for _, stat := range r.Stats {
		pct := float64(stat.Files) / float64(r.TotalFiles) * 100
		fmt.Printf("%-20s %8d %10s %7.1f%%\n", stat.Ext, stat.Files, formatBytes(stat.Bytes), pct)
	}

	fmt.Printf("%-20s %8s %10s %8s\n", "─────────────────", "─────────", "──────────", "───────")
	fmt.Printf("%-20s %8d %10s %7.1f%%\n", "Total", r.TotalFiles, formatBytes(r.TotalBytes), 100.0)
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
