.PHONY: build

all: build

build:
	go build -o elastic-search-archiver cmd/elastic-search-archiver/main.go
