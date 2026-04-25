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
filestats                        # analyze current directory
filestats /path/to/project       # analyze a specific path
filestats -l /path/to/project    # group by language
filestats -json /path/to/project # output as JSON
filestats -o stats.json /path    # save JSON to file, print table
filestats -version               # print version
```

Flags compose freely:

```bash
filestats -l -json              # language-grouped JSON to stdout
filestats -l -o stats.json      # language-grouped table + save JSON
```

## Options

| Flag         | Description                                                                 |
| ------------ | --------------------------------------------------------------------------- |
| `-l`         | Group results by language; Extension(s) column shows a comma-separated list |
| `-json`      | Print results as JSON to stdout instead of table                            |
| `-o file`    | Save results as JSON to a file (table still printed to stdout)              |
| `-version`   | Print version and exit                                                      |

## Releasing a new version

The release script keeps all versions in sync automatically (formula, git tag, binary ldflags):

```bash
make release V=1.2.3
```

This will:
1. Build binaries for all platforms into `dist/`
2. Compute SHA256 for each binary
3. Regenerate `Formula/filestats.rb` with the correct version and hashes
4. Commit the formula and create the git tag `v1.2.3`

Then follow the printed instructions to push and publish:

```bash
git push origin main v1.2.3
gh release create v1.2.3 dist/filestats-* --title "v1.2.3"
# copy Formula/filestats.rb to highnamee/homebrew-filestats and push
```

## Example output

Default (grouped by extension):

```
Extension   Language      Files        Size    Share
──────────  ────────  ─────────  ──────────  ───────  ───────────────
.go         Go                7     12.1 KB    53.8%  █████████░░░░░░
.gitignore  Git               1        19 B     7.7%  █░░░░░░░░░░░░░░
.md         Markdown          1      3.2 KB     7.7%  █░░░░░░░░░░░░░░
.mod        Go                1       117 B     7.7%  █░░░░░░░░░░░░░░
.sum        Go                1       812 B     7.7%  █░░░░░░░░░░░░░░
.yml        YAML              1       150 B     7.7%  █░░░░░░░░░░░░░░
Makefile    Make              1       204 B     7.7%  █░░░░░░░░░░░░░░
──────────  ────────  ─────────  ──────────  ───────  ───────────────
Total                        13     16.6 KB   100.0%  ███████████████
Completed in 3ms
```

Grouped by language (`-l`):

```
Language  Extension(s)         Files        Size    Share
────────  ───────────────  ─────────  ──────────  ───────  ───────────────
Go        .go, .mod, .sum          9     13.1 KB    69.2%  ███████████░░░░
Git       .gitignore               1        19 B     7.7%  █░░░░░░░░░░░░░░
Markdown  .md                      1      3.2 KB     7.7%  █░░░░░░░░░░░░░░
YAML      .yml                     1       150 B     7.7%  █░░░░░░░░░░░░░░
Make      Makefile                 1       204 B     7.7%  █░░░░░░░░░░░░░░
────────  ───────────────  ─────────  ──────────  ───────  ───────────────
Total                             13     16.8 KB   100.0%  ███████████████
Completed in 2ms
```

Results are sorted by file count descending. The `.git` directory and any paths matched by `.gitignore` are excluded.
