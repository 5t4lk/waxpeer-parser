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

	for k, v := range items {
		var price float64 = float64(items[k].Min)
		var itemPrice = price / 1000.00

		fmt.Printf("GUN: %s\nPRICE: %.2f$\n", v.Name, itemPrice)
	}
}
