.PHONY: test lint race

TEST_FLAGS?=-v
PKG_LIST := $(shell go list ./... | sort -u)

all: test

lint:
	golint -set_exit_status ${PKG_LIST}

race:
	go test -race -short ${PKG_LIST}

test:
	PATH="${PWD}/bin:${PWD}/test/bin:${PATH}" go test ${TEST_FLAGS} ${PKG_LIST}