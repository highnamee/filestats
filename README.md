# filestats

Count file statistics by extension, similar to GitHub's language breakdown.

## Installation

### Homebrew

```bash
brew tap highnamee/filestats
brew install filestats
```

### Build from source

```bash
go install github.com/highnamee/filestats@latest
```

## Requirements (development)

- Go 1.21+
- [golangci-lint](https://golangci-lint.run/welcome/install/) (for linting)

## Development

```bash
make build                          # compile binary to ./filestats
make run                            # run against current directory
make run ARGS="."                   # run against a specific path
make run ARGS="-l /path/to/project" # group by language, specific path
make lint                           # run golangci-lint
make fmt                            # format all Go files
make test                           # run tests
make clean                          # remove compiled binary
make release                        # cross-compile to dist/ for all platforms
```

## Run

```bash
filestats                                       # analyze current directory
filestats /path/to/project                      # analyze a specific path
filestats -l /path/to/project                   # group by language
filestats -json /path/to/project                # output as JSON
filestats -o stats.json /path                   # save JSON to file, print table
filestats -version                              # print version
filestats -exclude vendor                       # skip vendor directory
filestats -exclude vendor,node_modules          # skip multiple patterns
filestats -exclude vendor -exclude "*.min.js"   # repeatable flag form
filestats -loc=false                            # skip line counting (faster)
filestats -respect-gitignore=false              # include files ignored by .gitignore
```

Flags compose freely:

```bash
filestats -l -json                    # language-grouped JSON to stdout
filestats -l -o stats.json            # language-grouped table + save JSON
filestats -l -exclude vendor          # language-grouped, vendor excluded
filestats -loc=false -exclude vendor  # fast run, no LOC column
```

## Options

| Flag                       | Description                                                                 |
| -------------------------- | --------------------------------------------------------------------------- |
| `-l`                       | Group results by language; Extension(s) column shows a comma-separated list |
| `-top N`                   | Show only the top N results; remaining entries are aggregated into Others   |
| `-exclude pattern`         | Exclude files/dirs matching a glob pattern (repeatable, comma-separated)    |
| `-loc=false`               | Disable line counting; hides the Lines column and speeds up large repos     |
| `-respect-gitignore=false` | Include files that would normally be excluded by the root `.gitignore`      |
| `-json`                    | Print results as JSON to stdout instead of table                            |
| `-o file`                  | Save results as JSON to a file (table still printed to stdout)              |
| `-version`                 | Print version and exit                                                      |

## Example output

Default (grouped by extension):

```
Extension   Language      Files       Lines        Size    Share
──────────  ────────  ─────────  ──────────  ──────────  ───────  ───────────────
.go         Go                7         420     12.1 KB    53.8%  █████████░░░░░░
.gitignore  Git               1           2        19 B     7.7%  █░░░░░░░░░░░░░░
.md         Markdown          1          87      3.2 KB     7.7%  █░░░░░░░░░░░░░░
.mod        Go                1          10       117 B     7.7%  █░░░░░░░░░░░░░░
.sum        Go                1          21       812 B     7.7%  █░░░░░░░░░░░░░░
.yml        YAML              1          12       150 B     7.7%  █░░░░░░░░░░░░░░
Makefile    Make              1          20       204 B     7.7%  █░░░░░░░░░░░░░░
──────────  ────────  ─────────  ──────────  ──────────  ───────  ───────────────
Total                        13         572     16.6 KB   100.0%  ███████████████

--- Configuration ---
  Path          .
  Group by      extension
  Output        table
  Limit         all rows
  JSON file     (none)
  Excludes      .git/, root .gitignore
  Count LOC     yes
---------------------
Completed in 4ms
```

Grouped by language (`-l`):

```
Language  Extension(s)         Files       Lines        Size    Share
────────  ───────────────  ─────────  ──────────  ──────────  ───────  ───────────────
Go        .go, .mod, .sum          9         451     13.1 KB    69.2%  ███████████░░░░
Git       .gitignore               1           2        19 B     7.7%  █░░░░░░░░░░░░░░
Markdown  .md                      1          87      3.2 KB     7.7%  █░░░░░░░░░░░░░░
YAML      .yml                     1          12       150 B     7.7%  █░░░░░░░░░░░░░░
Make      Makefile                 1          20       204 B     7.7%  █░░░░░░░░░░░░░░
────────  ───────────────  ─────────  ──────────  ──────────  ───────  ───────────────
Total                             13         572     16.8 KB   100.0%  ███████████████

--- Configuration ---
  Path          .
  Group by      language
  Output        table
  Limit         all rows
  JSON file     (none)
  Excludes      .git/, root .gitignore
  Count LOC     yes
---------------------
Completed in 4ms
```

Results are sorted by file count descending. The `.git` directory and any paths matched by the root `.gitignore` are excluded by default (disable with `-respect-gitignore=false`).
