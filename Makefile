BIN := filestats

.PHONY: build run lint fmt test clean

build:
	go build -o $(BIN) .

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
