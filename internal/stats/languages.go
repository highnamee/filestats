package stats

import (
	"path/filepath"
	"strings"
)

// extensionLanguages maps lowercase file extensions to their friendly language name.
var extensionLanguages = map[string]string{
	// Go
	".go":  "Go",
	".mod": "Go",
	".sum": "Go",

	// Ruby
	".rb":       "Ruby",
	".ruby":     "Ruby",
	".rake":     "Ruby",
	".gemspec":  "Ruby",
	".ru":       "Ruby",
	".builder":  "Ruby",
	".jbuilder": "Ruby",
	".erb":      "ERB",
	".haml":     "Haml",
	".slim":     "Slim",

	// JavaScript / TypeScript
	".js":  "JavaScript",
	".jsx": "JavaScript",
	".mjs": "JavaScript",
	".cjs": "JavaScript",
	".ts":  "TypeScript",
	".tsx": "TypeScript",
	".ejs": "EJS",

	// Python
	".py":  "Python",
	".pyi": "Python",
	".pyw": "Python",

	// Rust
	".rs": "Rust",

	// Java / JVM
	".java":   "Java",
	".kt":     "Kotlin",
	".kts":    "Kotlin",
	".scala":  "Scala",
	".groovy": "Groovy",
	".gradle": "Gradle",

	// C / C++
	".c":   "C",
	".h":   "C/C++ Header",
	".cpp": "C++",
	".cc":  "C++",
	".cxx": "C++",
	".hpp": "C++ Header",

	// C# / .NET
	".cs":     "C#",
	".fs":     "F#",
	".fsx":    "F#",
	".vb":     "Visual Basic",
	".csproj": "MSBuild",
	".sln":    "Visual Studio",

	// Mobile
	".swift": "Swift",
	".dart":  "Dart",
	".m":     "Objective-C",
	".mm":    "Objective-C++",

	// PHP
	".php": "PHP",

	// Shell
	".sh":   "Shell",
	".bash": "Shell",
	".zsh":  "Shell",
	".fish": "Shell",
	".ps1":  "PowerShell",

	// Web
	".html":   "HTML",
	".htm":    "HTML",
	".css":    "CSS",
	".scss":   "SCSS",
	".sass":   "Sass",
	".less":   "Less",
	".vue":    "Vue",
	".svelte": "Svelte",

	// Data / Config
	".json":       "JSON",
	".json5":      "JSON5",
	".yaml":       "YAML",
	".yml":        "YAML",
	".toml":       "TOML",
	".xml":        "XML",
	".xsd":        "XML Schema",
	".xbrl":       "XBRL",
	".csv":        "CSV",
	".tsv":        "TSV",
	".sql":        "SQL",
	".properties": "Properties",
	".cnf":        "Config",
	".ini":        "Config",
	".env":        "Config",

	// Docs / Text
	".md":  "Markdown",
	".mdx": "MDX",
	".txt": "Text",
	".rst": "reStructuredText",
	".log": "Log",

	// Infrastructure
	".tf":      "Terraform",
	".tfvars":  "Terraform",
	".proto":   "Protobuf",
	".graphql": "GraphQL",
	".gql":     "GraphQL",

	// Images
	".png":  "Image",
	".jpg":  "Image",
	".jpeg": "Image",
	".gif":  "Image",
	".ico":  "Image",
	".svg":  "SVG",
	".webp": "Image",
	".bmp":  "Image",
	".cur":  "Image",

	// Fonts
	".ttf":   "Font",
	".otf":   "Font",
	".woff":  "Font",
	".woff2": "Font",
	".eot":   "Font",

	// Documents
	".pdf":  "PDF",
	".xlsx": "Excel",
	".xls":  "Excel",
	".ods":  "Spreadsheet",
	".docx": "Word",
	".doc":  "Word",

	// Lock files
	".lock": "Lock",
	".snap": "Snapshot",

	// Functional / other languages
	".lua":  "Lua",
	".r":    "R",
	".ex":   "Elixir",
	".exs":  "Elixir",
	".erl":  "Erlang",
	".hrl":  "Erlang",
	".hs":   "Haskell",
	".elm":  "Elm",
	".clj":  "Clojure",
	".cljs": "ClojureScript",
	".ml":   "OCaml",
	".mli":  "OCaml",
	".m4":   "M4",

	// Certs / Keys
	".der": "Certificate",
	".pem": "Certificate",
	".pub": "Public Key",

	// Git / VCS
	".gitignore":      "Git",
	".gitattributes":  "Git",
	".gitkeep":        "Git",
	".keep":           "Git",
	".git-pr-release": "Git",

	// Tooling dotfiles — these have no base name so filepath.Ext returns the whole name
	".npmrc":                    "npm",
	".nvmrc":                    "Node.js",
	".node-version":             "Node.js",
	".ruby-version":             "Ruby",
	".rspec":                    "Ruby",
	".prettierrc":               "Prettier",
	".prettierignore":           "Prettier",
	".eslintrc":                 "ESLint",
	".eslintignore":             "ESLint",
	".browserslistrc":           "Browserslist",
	".babelrc":                  "Babel",
	".editorconfig":             "EditorConfig",
	".dockerignore":             "Docker",
	".markuplintrc":             "MarkupLint",
	".yamllintignore":           "YAML",
	".openapi-generator-ignore": "OpenAPI",
	".db":                       "Database",
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
