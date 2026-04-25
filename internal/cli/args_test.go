package cli

import (
	"reflect"
	"testing"
)

func TestReorderArgs(t *testing.T) {
	bools := map[string]bool{"l": true, "json": true, "version": true}

	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "flags already before positional",
			in:   []string{"filestats", "-l", "-top", "5", "/path"},
			want: []string{"filestats", "-l", "-top", "5", "/path"},
		},
		{
			name: "positional before flags",
			in:   []string{"filestats", "/path", "-l", "-top", "5"},
			want: []string{"filestats", "-l", "-top", "5", "/path"},
		},
		{
			name: "bool flag after positional",
			in:   []string{"filestats", "/path", "-json"},
			want: []string{"filestats", "-json", "/path"},
		},
		{
			name: "value flag with inline =",
			in:   []string{"filestats", "-top=5", "/path"},
			want: []string{"filestats", "-top=5", "/path"},
		},
		{
			name: "no flags",
			in:   []string{"filestats", "/path"},
			want: []string{"filestats", "/path"},
		},
		{
			name: "no positional",
			in:   []string{"filestats", "-l", "-json"},
			want: []string{"filestats", "-l", "-json"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReorderArgs(tt.in, bools)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringsFlag(t *testing.T) {
	t.Run("single value", func(t *testing.T) {
		var f StringsFlag
		if err := f.Set("vendor"); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual([]string(f), []string{"vendor"}) {
			t.Errorf("got %v", f)
		}
	})

	t.Run("comma-separated", func(t *testing.T) {
		var f StringsFlag
		if err := f.Set("vendor,node_modules"); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual([]string(f), []string{"vendor", "node_modules"}) {
			t.Errorf("got %v", f)
		}
	})

	t.Run("repeated calls accumulate", func(t *testing.T) {
		var f StringsFlag
		_ = f.Set("vendor")
		_ = f.Set("dist")
		if !reflect.DeepEqual([]string(f), []string{"vendor", "dist"}) {
			t.Errorf("got %v", f)
		}
	})

	t.Run("String round-trips", func(t *testing.T) {
		var f StringsFlag
		_ = f.Set("a,b")
		if f.String() != "a,b" {
			t.Errorf("got %q", f.String())
		}
	})
}
