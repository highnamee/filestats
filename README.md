# filestats

Count file statistics by extension, similar to GitHub's language breakdown.

## Requirements

- Go 1.21+
- [golangci-lint](https://golangci-lint.run/welcome/install/) (for linting)

## Development

```bash
make build        # compile binary to ./filestats
make run          # run against current directory
make run ARGS=.   # run against a specific path
make lint         # run golangci-lint
make fmt          # format all Go files
make test         # run tests
make clean        # remove compiled binary
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

Or using the built binary:

```bash
./filestats
./filestats /path/to/project
```

## Example output

```
Extension               Files       Size    Share
─────────────────    ───────── ──────────  ───────
.go                         4     2.6 KB    57.1%
.json                       1      225 B    14.3%
.mod                        1       28 B    14.3%
─────────────────    ───────── ──────────  ───────
Total                       6     2.9 KB   100.0%
```

Results are sorted by file count descending. The `.git` directory and any paths matched by `.gitignore` are excluded.
