package filters

import (
	"fmt"
	"github.com/5t4lk/waxpeer-test/internal/apis/csgobackpack"
	"github.com/5t4lk/waxpeer-test/internal/apis/waxpeer"
	"strings"
)

func Filter() error {
	getByteJsonWX, err := waxpeer.GetRequest()
	if err != nil {
		return err
	}

	getWX, err := waxpeer.GetOrderedOutput(getByteJsonWX)
	if err != nil {
		return err
	}

	getByteJsonBP, err := csgobackpack.BPRequest()
	if err != nil {
		return err
	}

	getBP, err := csgobackpack.BPOrderedOutput(getByteJsonBP)
	if err != nil {
		return err
	}

	for _, v := range getWX {
		vSteamPrice := float64(v.SteamPrice) / 1000.00
		vPriceWax := float64(v.Min) / 1000.00

		diffPriceWaxpeer := vPriceWax / vSteamPrice
		if diffPriceWaxpeer <= 0.8 {
			for _, v2 := range getBP.ItemsList {
				if v.Name == v2.Name {
					diffMarketPrice := vSteamPrice / v2.Price.The30_Days.Median
					if diffMarketPrice >= 0.8 && diffMarketPrice <= 1.2 {
						if strings.Contains(v.Name, "Souvenir") {
							if diffPriceWaxpeer <= 0.3 {
								fmt.Printf("%s : %.2f : %.2f : %.1f%%\n", v.Name, vPriceWax, vSteamPrice, diffPriceWaxpeer*100)
							}
							break
						}
						if strings.Contains(v.Name, "★ StatTrak™") {
							if diffPriceWaxpeer <= 0.3 {
								fmt.Printf("%s : %.2f : %.2f : %.1f%%\n", v.Name, vPriceWax, vSteamPrice, diffPriceWaxpeer*100)
							}
							break
						}
						if strings.Contains(v.Name, "Doppler") {
							if diffPriceWaxpeer <= 0.3 {
								fmt.Printf("%s : %.2f : %.2f : %.1f%%\n", v.Name, vPriceWax, vSteamPrice, diffPriceWaxpeer*100)
							}
							break
						}
						if strings.Contains(v.Name, "Sticker") {
							if diffPriceWaxpeer <= 0.3 {
								fmt.Printf("%s : %.2f : %.2f : %.1f%%\n", v.Name, vPriceWax, vSteamPrice, diffPriceWaxpeer*100)
							}
							break
						}
						fmt.Printf("%s : %.2f : %.2f : %.1f%%\n", v.Name, vPriceWax, vSteamPrice, diffPriceWaxpeer*100)
					}
				}
			}
		}
	}
	fmt.Print("...\n")
	return nil
}
