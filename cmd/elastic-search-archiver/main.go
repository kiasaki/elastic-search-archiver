package main

import (
	esa "github.com/kiasaki/elastic-search-archiver"
)

func main() {
	config := esa.NewConfigFromFlags()
	instance := esa.New(config)
	instance.Run()
}
