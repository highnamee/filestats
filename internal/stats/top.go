package stats

// TopN returns a copy of r with Stats truncated to the n highest-ranked entries.
// A synthetic "Others" row aggregates the hidden entries so the table remains
// accurate. Totals are preserved so percentages reflect the full tree.
// If n <= 0 or n >= len(r.Stats) the result is returned unchanged.
func TopN(r *Result, n int) *Result {
	if n <= 0 || n >= len(r.Stats) {
		return r
	}

	others := ExtStat{Ext: "Others", Language: ""}
	for _, s := range r.Stats[n:] {
		others.Files += s.Files
		others.Bytes += s.Bytes
	}

	return &Result{
		Stats:             append(append([]ExtStat{}, r.Stats[:n]...), others),
		TotalFiles:        r.TotalFiles,
		TotalBytes:        r.TotalBytes,
		GroupedByLanguage: r.GroupedByLanguage,
		Trimmed:           true,
		Top:               n,
	}
}
