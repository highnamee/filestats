package stats

import (
	"fmt"
	"math"
	"strings"

	"github.com/fatih/color"
)

// barWidth is the number of characters used to render the share bar.
const barWidth = 15

var (
	styleHeader    = color.New(color.Bold)
	styleTotal     = color.New(color.Bold)
	styleDim       = color.New(color.Faint)
	styleGreenBold = color.New(color.FgGreen, color.Bold)
	styleGreen     = color.New(color.FgGreen)
	styleYellow    = color.New(color.FgYellow)
)

// languageColorMap maps language names to their display color.
var languageColorMap = map[string]*color.Color{
	"Go":         color.New(color.FgCyan),
	"Python":     color.New(color.FgYellow),
	"Ruby":       color.New(color.FgRed),
	"JavaScript": color.New(color.FgYellow),
	"TypeScript": color.New(color.FgBlue, color.Bold),
	"Rust":       color.New(color.FgRed),
	"Java":       color.New(color.FgRed),
	"Kotlin":     color.New(color.FgMagenta),
	"Scala":      color.New(color.FgRed),
	"Swift":      color.New(color.FgRed),
	"C":          color.New(color.FgBlue),
	"C++":        color.New(color.FgBlue),
	"C#":         color.New(color.FgMagenta),
	"F#":         color.New(color.FgMagenta),
	"PHP":        color.New(color.FgMagenta),
	"Dart":       color.New(color.FgBlue),
	"Elixir":     color.New(color.FgMagenta),
	"Erlang":     color.New(color.FgRed),
	"Haskell":    color.New(color.FgMagenta),
	"Clojure":    color.New(color.FgGreen),
	"OCaml":      color.New(color.FgYellow),
	"Elm":        color.New(color.FgCyan),
	"Lua":        color.New(color.FgBlue),
	"R":          color.New(color.FgBlue),
	"Shell":      color.New(color.FgGreen),
	"PowerShell": color.New(color.FgBlue),
	"HTML":       color.New(color.FgRed),
	"CSS":        color.New(color.FgBlue),
	"SCSS":       color.New(color.FgMagenta),
	"Sass":       color.New(color.FgMagenta),
	"Less":       color.New(color.FgBlue),
	"Vue":        color.New(color.FgGreen),
	"Svelte":     color.New(color.FgRed),
	"ERB":        color.New(color.FgRed),
	"Haml":       color.New(color.FgRed),
	"Slim":       color.New(color.FgRed),
	"EJS":        color.New(color.FgYellow),
	"JSON":       color.New(color.FgYellow),
	"YAML":       color.New(color.FgYellow),
	"TOML":       color.New(color.FgYellow),
	"XML":        color.New(color.FgYellow),
	"SQL":        color.New(color.FgBlue),
	"Markdown":   color.New(color.FgWhite),
	"CSV":        color.New(color.FgGreen),
	"Terraform":  color.New(color.FgMagenta),
	"Docker":     color.New(color.FgBlue),
	"Make":       color.New(color.FgGreen),
	"Git":        color.New(color.FgRed),
	"GraphQL":    color.New(color.FgMagenta),
	"Protobuf":   color.New(color.FgBlue),
	"Image":      color.New(color.FgCyan),
	"SVG":        color.New(color.FgCyan),
	"Font":       color.New(color.FgWhite),
	"PDF":        color.New(color.FgRed),
}

// coloredLang returns the language name styled with its associated color.
func coloredLang(lang string) string {
	if c, ok := languageColorMap[lang]; ok {
		return c.Sprint(lang)
	}
	return lang
}

// coloredShare returns the percentage string styled by magnitude.
func coloredShare(pct float64) string {
	s := fmt.Sprintf("%.1f%%", pct)
	switch {
	case pct >= 20:
		return styleGreenBold.Sprint(s)
	case pct >= 5:
		return styleGreen.Sprint(s)
	case pct >= 1:
		return styleYellow.Sprint(s)
	default:
		return styleDim.Sprint(s)
	}
}

// percentBar renders a fixed-width block bar for the given percentage.
func percentBar(pct float64) string {
	filled := int(math.Round(pct / 100 * barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)
	switch {
	case pct >= 20:
		return styleGreenBold.Sprint(bar)
	case pct >= 5:
		return styleGreen.Sprint(bar)
	case pct >= 1:
		return styleYellow.Sprint(bar)
	default:
		return styleDim.Sprint(bar)
	}
}

// padRight left-aligns s (visible length visLen) within a field of width w.
func padRight(s string, visLen, w int) string {
	if w > visLen {
		return s + strings.Repeat(" ", w-visLen)
	}
	return s
}

// padLeft right-aligns s (visible length visLen) within a field of width w.
func padLeft(s string, visLen, w int) string {
	if w > visLen {
		return strings.Repeat(" ", w-visLen) + s
	}
	return s
}
