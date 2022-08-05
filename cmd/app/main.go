package main

import (
	"github.com/5t4lk/waxpeer-test/internal/filters"
	"log"
)

func main() {
	for {
		if err := filters.Filter(); err != nil {
			log.Fatal(err)
		}
	}
}
