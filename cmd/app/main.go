package main

import (
	"github.com/5t4lk/waxpeer-test/internal/filters"
	"log"
)

func main() {
	for {
		err := filters.Filter()
		if err != nil {
			log.Fatal(err)
		}
	}
}
