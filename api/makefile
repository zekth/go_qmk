VERSION=$(shell git describe --always --tags)
BINARY:=dist/go_qmk

ifeq ($(OS), Windows_NT)
	BINARY:=$(BINARY).exe
	DETECT_RACE=''
endif

.PHONY: build
build:
	go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)" ./src

.PHONY: ci-build
ci-build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o $(BINARY) -ldflags "-s -w -X main.Version=$(VERSION)" ./src

.PHONY: test
test:
	go test -cover ./...
