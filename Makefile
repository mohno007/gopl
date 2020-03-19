.PHONY: test install install-run setup

export GO111MODULE := on

all: test

install:
	go install

install-run: install
	$(shell go env GOPATH)/bin/gopl

setup:
	brew install pre-commit
	brew install golangci/tap/golangci-lint
	pre-commit install
	export GO111MODULE=off && \
		go get -v github.com/fzipp/gocyclo && \
		go get -v github.com/go-lintpack/lintpack/... && \
		go get -v github.com/go-critic/go-critic/... && \
			cd "$(shell go env GOPATH)/src/github.com/go-critic/go-critic" && \
			make gocritic && \
			echo 'export PATH=$$PATH:$(shell go env GOPATH)/src/github.com/go-critic/go-critic'

test:
	go test ./...

test_verbose:
	go test -v ./...

cover:
	go test -cover ./...

cover_verbose: cover.out
	go tool cover -func=cover.out

cover_html: cover.html

cover.out: $(SRC)
	go test -coverprofile=cover.out ./...

cover.html: cover.out
	go tool cover -html=cover.out -o cover.html
