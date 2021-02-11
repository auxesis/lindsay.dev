.PHONY: test

all: build

build:
	go build server.go

#export GOOS := linux
export GOARCH := amd64
