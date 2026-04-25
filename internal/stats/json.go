package stats

import (
	"encoding/json"
	"io"
	"os"
)

// WriteJSON encodes r as indented JSON and writes it to w.
func WriteJSON(w io.Writer, r *Result) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(r)
}

// SaveJSON writes r as indented JSON to the file at path, creating or truncating it.
func SaveJSON(path string, r *Result) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := WriteJSON(f, r); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}
