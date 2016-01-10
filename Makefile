GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@go build $(GOFLAGS) ./...

format:
	@gofmt -l -s -w .

lint:
	@golint .

optimize_imports:
	@goimports -l -w .

beautify: format lint optimize_imports

vet:
	@go vet ./...

race:
	@go test -race $(GOFLAGS) ./...

audit: vet race

install:
	@go get $(GOFLAGS) ./...

test: install
	@go test $(GOFLAGS) ./...

checkin : beautify audit
	@git add -p .

bench: install
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

serve: build install
	@chat-server -c config.json