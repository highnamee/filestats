package stats

// ExtStat holds the aggregated file count and total size for a single extension or special filename.
type ExtStat struct {
	Ext      string `json:"ext"`
	Language string `json:"language,omitempty"`
	Files    int    `json:"files"`
	Lines    int64  `json:"lines"`
	Bytes    int64  `json:"bytes"`
}

// Result holds the full analysis output for a directory tree.
type Result struct {
	Stats             []ExtStat `json:"stats"`
	TotalFiles        int       `json:"total_files"`
	TotalLines        int64     `json:"total_lines"`
	TotalBytes        int64     `json:"total_bytes"`
	GroupedByLanguage bool      `json:"grouped_by_language"`
	Trimmed           bool      `json:"trimmed,omitempty"`
	Top               int       `json:"top,omitempty"`
}
