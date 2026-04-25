# filestats

Count file statistics by extension, similar to GitHub's language breakdown.

## Requirements

- Go 1.21+
- [golangci-lint](https://golangci-lint.run/welcome/install/) (for linting)

## Development

```bash
make build # compile binary to ./filestats
make run # run against current directory
make run ARGS="." # run against a specific path
make run ARGS="-l" # group by language, current directory
make run ARGS="-l /path/to/project" # group by language, specific path
make lint # run golangci-lint
make fmt # format all Go files
make test # run tests
make clean # remove compiled binary
```

## Run

Analyze the current directory:

```bash
go run .
```

Analyze a specific path:

```bash
go run . /path/to/project
```

Group results by language instead of extension:

```bash
go run . -l
go run . -l /path/to/project
```

Output as JSON:

```bash
go run . -json
go run . -json /path/to/project
```

Save JSON to a file (table is still printed to stdout):

```bash
go run . -o stats.json
go run . -o stats.json /path/to/project
```

Flags compose freely:

```bash
go run . -l -json              # language-grouped JSON to stdout
go run . -l -o stats.json      # language-grouped table + save JSON
go run . -l -json -o stats.json # language-grouped JSON to stdout + save JSON
```

Or using the built binary:

```bash
./filestats
./filestats /path/to/project
./filestats -l /path/to/project
./filestats -json /path/to/project
./filestats -o stats.json /path/to/project
```

## Options

| Flag       | Description                                                                 |
| ---------- | --------------------------------------------------------------------------- |
| `-l`       | Group results by language; Extension(s) column shows a comma-separated list |
| `-json`    | Print results as JSON to stdout instead of table                            |
| `-o file`  | Save results as JSON to a file (table still printed to stdout)              |

## Example output

Default (grouped by extension):

```
Extension            Language            Files       Size    Share
──────────────────── ──────────────────  ──────── ──────────  ───────
.go                  Go                      7     8.8 KB    53.8%
.gitignore           Git                     1       19 B     7.7%
.md                  Markdown                1     1.4 KB     7.7%
.mod                 Go                      1      117 B     7.7%
.sum                 Go                      1      812 B     7.7%
.yml                 YAML                    1      150 B     7.7%
Makefile             Make                    1      204 B     7.7%
──────────────────── ──────────────────  ──────── ──────────  ───────
Total                                       13    11.5 KB   100.0%
```

Grouped by language (`-l`):

```
Language         Extension(s)            Files       Size    Share
──────────────── ────────────────────  ──────── ──────────  ───────
Go               .go, .mod, .sum             9     9.8 KB    69.2%
Git              .gitignore                  1       19 B     7.7%
Markdown         .md                         1     1.4 KB     7.7%
YAML             .yml                        1      150 B     7.7%
Make             Makefile                    1      204 B     7.7%
──────────────────── ──────────────────  ──────── ──────────  ───────
Total                                       13    11.5 KB   100.0%
```

Results are sorted by file count descending. The `.git` directory and any paths matched by `.gitignore` are excluded.
