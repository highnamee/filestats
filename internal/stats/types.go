package stats

// ExtStat holds the aggregated file count and total size for a single extension.
type ExtStat struct {
	Ext   string
	Files int
	Bytes int64
}

// Result holds the full analysis output for a directory tree.
type Result struct {
	Stats      []ExtStat
	TotalFiles int
	TotalBytes int64
}
