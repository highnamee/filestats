package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// PrintFlagDefaults writes flag defaults to w with "(default true)" highlighted in green.
func PrintFlagDefaults(w io.Writer) error {
	var buf strings.Builder
	flag.CommandLine.SetOutput(&buf)
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(w)
	green := color.New(color.FgGreen).SprintFunc()
	_, err := fmt.Fprint(w, strings.ReplaceAll(buf.String(), "(default true)", green("(default true)")))
	return err
}
