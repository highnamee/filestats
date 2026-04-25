package stats

import (
	"path/filepath"
	"strings"
)

// extensionLanguages maps lowercase file extensions to their friendly language name.
var extensionLanguages = map[string]string{
	".go":            "Go",
	".mod":           "Go",
	".sum":           "Go",
	".py":            "Python",
	".rb":            "Ruby",
	".js":            "JavaScript",
	".jsx":           "JavaScript",
	".ts":            "TypeScript",
	".tsx":           "TypeScript",
	".rs":            "Rust",
	".java":          "Java",
	".c":             "C",
	".h":             "C/C++ Header",
	".cpp":           "C++",
	".cc":            "C++",
	".cxx":           "C++",
	".hpp":           "C++ Header",
	".cs":            "C#",
	".php":           "PHP",
	".swift":         "Swift",
	".kt":            "Kotlin",
	".kts":           "Kotlin",
	".scala":         "Scala",
	".sh":            "Shell",
	".bash":          "Shell",
	".zsh":           "Shell",
	".fish":          "Shell",
	".html":          "HTML",
	".htm":           "HTML",
	".css":           "CSS",
	".scss":          "SCSS",
	".sass":          "Sass",
	".less":          "Less",
	".json":          "JSON",
	".yaml":          "YAML",
	".yml":           "YAML",
	".toml":          "TOML",
	".xml":           "XML",
	".sql":           "SQL",
	".md":            "Markdown",
	".mdx":           "MDX",
	".txt":           "Text",
	".gitignore":     "Git",
	".gitattributes": "Git",
	".csv":           "CSV",
	".tf":            "Terraform",
	".tfvars":        "Terraform",
	".proto":         "Protobuf",
	".graphql":       "GraphQL",
	".gql":           "GraphQL",
	".vue":           "Vue",
	".svelte":        "Svelte",
	".dart":          "Dart",
	".lua":           "Lua",
	".r":             "R",
	".ex":            "Elixir",
	".exs":           "Elixir",
	".erl":           "Erlang",
	".hrl":           "Erlang",
	".hs":            "Haskell",
	".elm":           "Elm",
	".clj":           "Clojure",
	".ml":            "OCaml",
	".mli":           "OCaml",
	".fs":            "F#",
	".fsx":           "F#",
}

// filenameLanguages maps exact filenames (no extension) to their friendly language name.
var filenameLanguages = map[string]string{
	"Makefile":    "Make",
	"makefile":    "Make",
	"GNUmakefile": "Make",
	"Dockerfile":  "Docker",
	"Gemfile":     "Ruby",
	"Rakefile":    "Ruby",
	"Guardfile":   "Ruby",
	"Vagrantfile": "Ruby",
	"Podfile":     "Ruby",
	"Fastfile":    "Fastlane",
	"Appfile":     "Fastlane",
	"Matchfile":   "Fastlane",
	"Brewfile":    "Homebrew",
	"Procfile":    "Procfile",
	"Jenkinsfile": "Jenkins",
}

// groupKey returns the grouping key for a filename: the lowercased extension if present,
// the exact filename if it is a recognised special file, or "(no extension)" otherwise.
func groupKey(name string) string {
	if ext := strings.ToLower(filepath.Ext(name)); ext != "" {
		return ext
	}
	if _, ok := filenameLanguages[name]; ok {
		return name
	}
	return "(no extension)"
}

// languageFor returns the friendly language name for a given key (extension or filename).
// Returns an empty string if unknown.
func languageFor(key string) string {
	if lang, ok := extensionLanguages[key]; ok {
		return lang
	}
	if lang, ok := filenameLanguages[key]; ok {
		return lang
	}
	return ""
}
