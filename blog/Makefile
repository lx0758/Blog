.PHONY: init dev test build build-linux build-windows clean

VERSION ?= $(shell git describe --tags --always 2>/dev/null || echo "dev")

init:
	go mod tidy

dev:
	go run main.go

test:
	go test ./...

build: build-linux build-windows

build-linux:
	echo "Building blog for linux/amd64"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=gcc go build -ldflags '-X blog/res.Version=$(VERSION) -linkmode external -extldflags "-static"' -o bin/blog-$(VERSION) -v main.go
	echo "Building blog for linux/amd64"

build-windows:
	echo "Building blog for windows/amd64"
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags '-X blog/res.Version=$(VERSION) -linkmode external -extldflags "-static"' -o bin/blog-$(VERSION).exe -v main.go
	echo "Building blog for windows/amd64"

clean:
	rm -rf bin

.DEFAULT_GOAL := build
