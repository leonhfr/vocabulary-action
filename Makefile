.PHONY: default
default: build

.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: doc
doc:
	godoc -http=:6060

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...
