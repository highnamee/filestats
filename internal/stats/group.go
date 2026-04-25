package stats

import (
	"sort"
	"strings"
)

// GroupByLanguage re-groups a Result by language name.
// Each entry's Ext becomes a comma-separated list of extensions that belong to that language.
// Entries with no recognised language are collected under "(unknown)".
func GroupByLanguage(r *Result) *Result {
	type bucket struct {
		exts  []string
		files int
		bytes int64
	}

	groups := make(map[string]*bucket)

	for _, stat := range r.Stats {
		lang := stat.Language
		if lang == "" {
			lang = "(unknown)"
		}
		if groups[lang] == nil {
			groups[lang] = &bucket{}
		}
		groups[lang].exts = append(groups[lang].exts, stat.Ext)
		groups[lang].files += stat.Files
		groups[lang].bytes += stat.Bytes
	}

	result := &Result{
		TotalFiles:        r.TotalFiles,
		TotalBytes:        r.TotalBytes,
		GroupedByLanguage: true,
	}

	for lang, b := range groups {
		sort.Strings(b.exts)
		result.Stats = append(result.Stats, ExtStat{
			Ext:      strings.Join(b.exts, ", "),
			Language: lang,
			Files:    b.files,
			Bytes:    b.bytes,
		})
	}

	sortByFiles(result.Stats)
	return result
}
