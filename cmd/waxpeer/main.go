package main

import (
	"444/internal/client"
	"fmt"
	"log"
)

func main() {
	getJson, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}

	items, err := client.OrderedOutput(getJson)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range items {
		fmt.Printf("GUN: %s\nPRICE: %d\n", v.Name, v.Min)
	}
}
