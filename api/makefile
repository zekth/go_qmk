VERSION=$(shell git describe --always --tags)
BINARY:=dist/go_qmk
ifeq ($(OS), Windows_NT)
	BINARY:=$(BINARY).exe
	DETECT_RACE=''
endif

# Getting mandatory component for dev env
.PHONY: init
init:
	go get github.com/cespare/reflex
	brew install golangci/tap/golangci-lint

.PHONY: build
build:
	go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY: ci-build
ci-build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint:
	golangci-lint run

# Hot reload
.PHONY: dev
dev:
	reflex -r '\.go' -s -- sh -c 'make run'

.PHONY: run
run:
	go run main.go
