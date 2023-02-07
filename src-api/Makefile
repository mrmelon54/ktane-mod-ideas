SHELL := /bin/bash
BIN := dist/ktane-mod-ideas
HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH}
LD_FLAGS := -s -w -X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'
COMP_BIN := go

.PHONY: build dev clean

build:
	mkdir -p dist/
	${COMP_BIN} build -o "${BIN}" -ldflags="${LD_FLAGS}" ./cmd/ktane-mod-ideas

dev:
	${COMP_BIN} build -tags debug -o "${BIN}" -ldflags="${LD_FLAGS}" ./cmd/ktane-mod-ideas
	./${BIN}

clean:
	go clean
	rm ${BIN}
