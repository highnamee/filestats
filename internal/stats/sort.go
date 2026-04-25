package stats

import "sort"

// sortByFiles sorts entries by file count descending, breaking ties alphabetically by extension.
func sortByFiles(entries []ExtStat) {
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Files != entries[j].Files {
			return entries[i].Files > entries[j].Files
		}
		return entries[i].Ext < entries[j].Ext
	})
}
