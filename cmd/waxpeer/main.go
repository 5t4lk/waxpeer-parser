package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	getByteJson, err := waxpeer.Request()
	if err != nil {
		log.Fatal(err)
	}

	items, err := waxpeer.OrderedOutput(getByteJson, "")
	if err != nil {
		log.Fatal(err)
	}

	clientShadowpay := shadowpay.New("a163b4af51907ad6df96f684cc3c67b2")

	jsonByteInfo, err := clientShadowpay.SPRequest()
	if err != nil {
		log.Fatal(err)
	}

	SPitems, err := shadowpay.SPOrderedOutput(jsonByteInfo)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range items {
		for _, v2 := range SPitems {
			if v.Name == v2.SteamItem.SteamMarketHashName {
				var temp = float64(v.Min)
				var waxItemPrice = temp / 1000.00
				var diffPrice = waxItemPrice / v2.SteamItem.SuggestedPrice
				if diffPrice < 0.7 {
					if strings.Contains(v.Name, "Sticker") {
						continue
					}
					if strings.Contains(v.Name, "StatTrak™") {
						continue
					}
					fmt.Printf("%s : %.2f : %.2f : %.2f\n", v.Name, waxItemPrice, v2.SteamItem.SuggestedPrice, diffPrice)
				}
			}
		}
	}
	/*	for _, itemS := range SPitems {
			//for k, _ := range items {
			//	var price float64 = float64(items[k].Min)
			//	var itemPrice = price / 1000.00
			//
			//	var diffPrice = itemPrice / itemS.SteamItem.SuggestedPrice
			//	fmt.Println(itemS.SteamItem.SteamMarketHashName, itemS.SteamItem.SuggestedPrice, diffPrice)
			//}
			fmt.Println(itemS.SteamItem.SuggestedPrice)
		}
		fmt.Println(items)*/
}
