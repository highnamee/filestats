BIN     := filestats
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: build run lint fmt test clean release

build:
	go build $(LDFLAGS) -o $(BIN) .

run:
	go run . $(ARGS)

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .

test:
	go test ./...

clean:
	rm -f $(BIN)

release:
	@test -n "$(V)" || (echo "usage: make release V=1.2.3"; exit 1)
	bash scripts/release.sh $(V)
