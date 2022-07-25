package main

import (
	"444/internal/apis/waxpeer"
	"fmt"
	"log"
)

func main() {
	getJson, err := waxpeer.Get()
	if err != nil {
		log.Fatal(err)
	}

	items, err := waxpeer.OrderedOutput(getJson)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range items {
		fmt.Printf("GUN: %s\nPRICE: %d\n", v.Name, v.Min)
	}
}
g