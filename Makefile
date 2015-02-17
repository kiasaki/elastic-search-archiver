.PHONY: build dist

all: build

build:
	go build -o elastic-search-archiver cmd/elastic-search-archiver/main.go

dist:
	GOOS=linux go build -o elastic-search-archiver-linux cmd/elastic-search-archiver/main.go
	GOOS=darwin go build -o elastic-search-archiver-darwin cmd/elastic-search-archiver/main.go
