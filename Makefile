.PHONY: init build build-admin build-blog clean

VERSION ?= $(shell git describe --tags --always 2>/dev/null || echo "dev")

all: build

init:
	cd admin && $(MAKE) init
	cd blog && $(MAKE) init

build: build-blog

build-admin:
	cd admin && VITE_APP_VERSION=$(VERSION) $(MAKE) build

build-blog: build-admin
	cd blog && VERSION=$(VERSION) $(MAKE) build

clean:
	cd admin && $(MAKE) clean
	cd blog && $(MAKE) clean

.DEFAULT_GOAL := build
