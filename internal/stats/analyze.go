// Package stats provides file-system analysis grouped by file extension.
package stats

import (
	"os"
	"path/filepath"
	"sync"

	ignore "github.com/sabhiram/go-gitignore"
)

type fileEntry struct {
	key  string
	size int64
}

// Analyze walks root recursively, counting files grouped by extension.
// Directories and files matched by any .gitignore found in root are excluded.
// The .git directory is always skipped. Directories are processed concurrently.
func Analyze(root string) (*Result, error) {
	gi, _ := ignore.CompileIgnoreFile(filepath.Join(root, ".gitignore"))

	results := make(chan fileEntry, 512)
	var wg sync.WaitGroup

	var walkDir func(dir, rel string)
	walkDir = func(dir, rel string) {
		defer wg.Done()

		entries, err := os.ReadDir(dir)
		if err != nil {
			return
		}

		for _, de := range entries {
			name := de.Name()
			entryRel := filepath.Join(rel, name)

			if de.IsDir() {
				if name == ".git" {
					continue
				}
				if gi != nil && gi.MatchesPath(entryRel) {
					continue
				}
				wg.Add(1)
				go walkDir(filepath.Join(dir, name), entryRel)
				continue
			}

			if gi != nil && gi.MatchesPath(entryRel) {
				continue
			}

			info, err := de.Info()
			if err != nil {
				continue
			}

			results <- fileEntry{key: groupKey(name), size: info.Size()}
		}
	}

	wg.Add(1)
	go walkDir(root, "")

	go func() {
		wg.Wait()
		close(results)
	}()

	counts := make(map[string]*ExtStat)
	for fe := range results {
		if counts[fe.key] == nil {
			counts[fe.key] = &ExtStat{Ext: fe.key, Language: languageFor(fe.key)}
		}
		counts[fe.key].Files++
		counts[fe.key].Bytes += fe.size
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
