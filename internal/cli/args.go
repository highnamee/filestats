// Package cli provides command-line argument utilities.
package cli

import "strings"

// StringsFlag is a flag.Value that accumulates repeated -flag values.
// It also accepts comma-separated values in a single flag invocation.
type StringsFlag []string

func (f *StringsFlag) String() string { return strings.Join(*f, ",") }

// Set appends one or more comma-separated values to the flag.
func (f *StringsFlag) Set(v string) error {
	*f = append(*f, strings.Split(v, ",")...)
	return nil
}

// ReorderArgs moves flag arguments before positional arguments so that
// flag.Parse can see them regardless of where the user placed them.
// boolFlags lists flag names that take no value (e.g. "l", "json").
func ReorderArgs(args []string, boolFlags map[string]bool) []string {
	out := args[:1:1] // keep program name
	var pos []string

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			pos = append(pos, arg)
			continue
		}
		// Extract flag name, stripping dashes and any =value suffix.
		name := strings.TrimLeft(arg, "-")
		if idx := strings.IndexByte(name, '='); idx >= 0 {
			name = name[:idx]
		}
		out = append(out, arg)
		// If this flag expects a value, consume the next arg as its value.
		if !boolFlags[name] && i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
			i++
			out = append(out, args[i])
		}
	}

	return append(out, pos...)
}
