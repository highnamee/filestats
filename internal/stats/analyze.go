// Package stats provides file-system analysis grouped by file extension.
package stats

import (
	"io/fs"
	"os"
	"path/filepath"

	ignore "github.com/sabhiram/go-gitignore"
)

// Analyze walks root recursively, counting files grouped by extension.
// Directories and files matched by any .gitignore found in root are excluded.
// The .git directory is always skipped.
func Analyze(root string) (*Result, error) {
	gi, _ := ignore.CompileIgnoreFile(filepath.Join(root, ".gitignore"))

	counts := make(map[string]*ExtStat)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return nil
		}

		if d.IsDir() {
			if d.Name() == ".git" {
				return filepath.SkipDir
			}
			if gi != nil && gi.MatchesPath(rel) {
				return filepath.SkipDir
			}
			return nil
		}

		if gi != nil && gi.MatchesPath(rel) {
			return nil
		}

		key := groupKey(d.Name())

		info, err := os.Stat(path)
		if err != nil {
			return nil
		}

		if counts[key] == nil {
			counts[key] = &ExtStat{Ext: key, Language: languageFor(key)}
		}
		counts[key].Files++
		counts[key].Bytes += info.Size()
		return nil
	})
	if err != nil {
		return nil, err
	}

	result := &Result{}
	for _, stat := range counts {
		result.Stats = append(result.Stats, *stat)
		result.TotalFiles += stat.Files
		result.TotalBytes += stat.Bytes
	}

	sortByFiles(result.Stats)
	return result, nil
}
