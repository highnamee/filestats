package stats

import (
	"strings"
	"testing"
)

func TestGroupByLanguage(t *testing.T) {
	r := &Result{
		TotalFiles: 6,
		TotalBytes: 600,
		Stats: []ExtStat{
			{Ext: ".go", Language: "Go", Files: 3, Bytes: 300},
			{Ext: ".mod", Language: "Go", Files: 1, Bytes: 100},
			{Ext: ".rb", Language: "Ruby", Files: 2, Bytes: 200},
		},
	}

	got := GroupByLanguage(r)

	if !got.GroupedByLanguage {
		t.Error("GroupedByLanguage should be true")
	}
	if got.TotalFiles != r.TotalFiles || got.TotalBytes != r.TotalBytes {
		t.Error("totals should be preserved")
	}
	if len(got.Stats) != 2 {
		t.Fatalf("expected 2 language groups, got %d", len(got.Stats))
	}

	// First row should be Go (most files).
	goRow := got.Stats[0]
	if goRow.Language != "Go" {
		t.Errorf("first row language = %q, want Go", goRow.Language)
	}
	if goRow.Files != 4 {
		t.Errorf("Go files = %d, want 4", goRow.Files)
	}
	// Extensions are sorted and joined.
	if !strings.Contains(goRow.Ext, ".go") || !strings.Contains(goRow.Ext, ".mod") {
		t.Errorf("Go extensions = %q, want both .go and .mod", goRow.Ext)
	}

	ruby := got.Stats[1]
	if ruby.Language != "Ruby" || ruby.Files != 2 {
		t.Errorf("second row = %+v, want Ruby/2 files", ruby)
	}
}

func TestGroupByLanguage_unknown(t *testing.T) {
	r := &Result{
		TotalFiles: 1,
		TotalBytes: 10,
		Stats: []ExtStat{
			{Ext: ".xyz", Language: "", Files: 1, Bytes: 10},
		},
	}

	got := GroupByLanguage(r)
	if len(got.Stats) != 1 || got.Stats[0].Language != "(unknown)" {
		t.Errorf("unrecognised extension should be grouped as (unknown), got %+v", got.Stats)
	}
}
