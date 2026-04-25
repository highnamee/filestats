// Package stats provides file-system analysis grouped by file extension.
package stats

import (
	"bytes"
	"os"
	"path/filepath"
	"sync"

	ignore "github.com/sabhiram/go-gitignore"
)

type fileEntry struct {
	key   string
	size  int64
	lines int64
}

// Analyze walks root recursively, counting files grouped by extension.
// Directories and files matched by any .gitignore found in root are excluded.
// The .git directory is always skipped. Directories are processed concurrently.
// Entries matching any pattern in excludes are also skipped (gitignore semantics).
func Analyze(root string, excludes []string, loc bool) (*Result, error) {
	gi, _ := ignore.CompileIgnoreFile(filepath.Join(root, ".gitignore"))
	excl := ignore.CompileIgnoreLines(excludes...)

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

			if excl.MatchesPath(entryRel) || (gi != nil && gi.MatchesPath(entryRel)) {
				continue
			}

			if de.IsDir() {
				if name == ".git" {
					continue
				}
				wg.Add(1)
				go walkDir(filepath.Join(dir, name), entryRel)
				continue
			}

			info, err := de.Info()
			if err != nil {
				continue
			}

			fe := fileEntry{key: groupKey(name), size: info.Size()}
			if loc {
				fe.lines = countLines(filepath.Join(dir, name))
			}
			results <- fe
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
		counts[fe.key].Lines += fe.lines
		counts[fe.key].Bytes += fe.size
	}

	result := &Result{}
	for _, stat := range counts {
		result.Stats = append(result.Stats, *stat)
		result.TotalFiles += stat.Files
		result.TotalLines += stat.Lines
		result.TotalBytes += stat.Bytes
	}

	sortByFiles(result.Stats)
	return result, nil
}

// countLines counts newlines in a file. Returns 0 for binary files (detected
// by a null byte in the first read chunk).
func countLines(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer func() { _ = f.Close() }()

	buf := make([]byte, 32*1024)
	var lines int64
	firstChunk := true
	var lastByte byte

	for {
		n, err := f.Read(buf)
		if n > 0 {
			chunk := buf[:n]
			if firstChunk {
				if bytes.IndexByte(chunk, 0) >= 0 {
					return 0
				}
				firstChunk = false
			}
			lines += int64(bytes.Count(chunk, []byte{'\n'}))
			lastByte = chunk[n-1]
		}
		if err != nil {
			break
		}
	}
	// File has content but the last line has no trailing newline.
	if !firstChunk && lastByte != '\n' {
		lines++
	}
	return lines
}
