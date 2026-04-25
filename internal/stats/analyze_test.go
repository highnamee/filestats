package stats

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestAnalyze_basic(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "main.go"), "package main")
	writeFile(t, filepath.Join(dir, "util.go"), "package main")
	writeFile(t, filepath.Join(dir, "README.md"), "# hello")
	writeFile(t, filepath.Join(dir, "sub", "helper.go"), "package sub")

	r, err := Analyze(dir, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	if r.TotalFiles != 4 {
		t.Errorf("TotalFiles = %d, want 4", r.TotalFiles)
	}
	if r.TotalLines == 0 {
		t.Error("TotalLines should be > 0")
	}

	counts := statsByExt(r)
	if counts[".go"] != 3 {
		t.Errorf(".go count = %d, want 3", counts[".go"])
	}
	if counts[".md"] != 1 {
		t.Errorf(".md count = %d, want 1", counts[".md"])
	}
}

func TestAnalyze_exclude_dir(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "main.go"), "package main")
	writeFile(t, filepath.Join(dir, "vendor", "lib.go"), "package lib")
	writeFile(t, filepath.Join(dir, "vendor", "dep.go"), "package dep")

	r, err := Analyze(dir, []string{"vendor"}, false)
	if err != nil {
		t.Fatal(err)
	}

	if r.TotalFiles != 1 {
		t.Errorf("TotalFiles = %d, want 1 (vendor should be excluded)", r.TotalFiles)
	}
}

func TestAnalyze_exclude_glob(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "app.go"), "package main")
	writeFile(t, filepath.Join(dir, "app.min.js"), "min")
	writeFile(t, filepath.Join(dir, "bundle.min.js"), "min")

	r, err := Analyze(dir, []string{"*.min.js"}, false)
	if err != nil {
		t.Fatal(err)
	}

	if r.TotalFiles != 1 {
		t.Errorf("TotalFiles = %d, want 1 (*.min.js should be excluded)", r.TotalFiles)
	}
}

func TestAnalyze_exclude_multiple(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "main.go"), "package main")
	writeFile(t, filepath.Join(dir, "vendor", "a.go"), "a")
	writeFile(t, filepath.Join(dir, "node_modules", "b.js"), "b")

	r, err := Analyze(dir, []string{"vendor", "node_modules"}, false)
	if err != nil {
		t.Fatal(err)
	}

	if r.TotalFiles != 1 {
		t.Errorf("TotalFiles = %d, want 1", r.TotalFiles)
	}
}

func TestAnalyze_gitignore(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, ".gitignore"), "dist/\n*.log\n")
	writeFile(t, filepath.Join(dir, "main.go"), "package main")
	writeFile(t, filepath.Join(dir, "dist", "out.js"), "built")
	writeFile(t, filepath.Join(dir, "debug.log"), "log")

	r, err := Analyze(dir, nil, true)
	if err != nil {
		t.Fatal(err)
	}

	// Only main.go should be counted; .gitignore itself is also counted.
	counts := statsByExt(r)
	if counts[".go"] != 1 {
		t.Errorf(".go count = %d, want 1", counts[".go"])
	}
	if counts[".js"] != 0 {
		t.Errorf(".js count = %d, want 0 (dist/ gitignored)", counts[".js"])
	}
	if counts[".log"] != 0 {
		t.Errorf(".log count = %d, want 0 (*.log gitignored)", counts[".log"])
	}
}

// statsByExt returns a map of extension → file count from a Result.
func statsByExt(r *Result) map[string]int {
	m := make(map[string]int, len(r.Stats))
	for _, s := range r.Stats {
		m[s.Ext] = s.Files
	}
	return m
}
