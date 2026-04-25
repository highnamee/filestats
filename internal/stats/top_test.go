package stats

import (
	"testing"
)

func makeResult(exts ...string) *Result {
	r := &Result{}
	for i, ext := range exts {
		r.Stats = append(r.Stats, ExtStat{Ext: ext, Files: len(exts) - i, Bytes: int64(100 * (len(exts) - i))})
		r.TotalFiles += len(exts) - i
		r.TotalBytes += int64(100 * (len(exts) - i))
	}
	return r
}

func TestTopN_noop(t *testing.T) {
	r := makeResult(".go", ".rb", ".js")

	if got := TopN(r, 0); got != r {
		t.Error("n=0 should return original result")
	}
	if got := TopN(r, 3); got != r {
		t.Error("n=len should return original result")
	}
	if got := TopN(r, 5); got != r {
		t.Error("n>len should return original result")
	}
}

func TestTopN_trim(t *testing.T) {
	r := makeResult(".go", ".rb", ".js", ".md")
	got := TopN(r, 2)

	if len(got.Stats) != 3 {
		t.Fatalf("expected 3 rows (top 2 + Others), got %d", len(got.Stats))
	}

	if got.Stats[2].Ext != "Others" {
		t.Errorf("last row should be Others, got %q", got.Stats[2].Ext)
	}

	// Totals must be preserved.
	if got.TotalFiles != r.TotalFiles {
		t.Errorf("TotalFiles changed: got %d, want %d", got.TotalFiles, r.TotalFiles)
	}
	if got.TotalBytes != r.TotalBytes {
		t.Errorf("TotalBytes changed: got %d, want %d", got.TotalBytes, r.TotalBytes)
	}

	// Others aggregates the hidden rows.
	wantOthersFiles := r.Stats[2].Files + r.Stats[3].Files
	if got.Stats[2].Files != wantOthersFiles {
		t.Errorf("Others.Files = %d, want %d", got.Stats[2].Files, wantOthersFiles)
	}

	if !got.Trimmed || got.Top != 2 {
		t.Errorf("Trimmed/Top not set correctly: Trimmed=%v Top=%d", got.Trimmed, got.Top)
	}
}
