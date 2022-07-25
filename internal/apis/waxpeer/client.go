package waxpeer

import (
	"444/internal/entities"
	"444/internal/special_item"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	minPriceURL = "https://api.waxpeer.com/v1/prices?game=csgo&min_price="
	maxPriceURL = "&max_price="
	searchURL   = "&search="
)

func Request(min, max, search string) ([]byte, error) {
	reqURL := minPriceURL + min + maxPriceURL + max + searchURL + search

	r, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	byteInfo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteInfo, nil
}

func Get() ([]byte, error) {
	minPrice, maxPrice, itemName, err := special_item.ItemInfo()
	if err != nil {
		return nil, err
	}

	byteJson, err := Request(minPrice, maxPrice, itemName)
	if err != nil {
		return nil, err
	}

	return byteJson, nil
}

func OrderedOutput(input []byte) ([]entities.Items, error) {
	var items entities.ItemsData

	if err := json.Unmarshal(input, &items); err != nil {
		return nil, err
	}

	if !items.Success {
		return nil, fmt.Errorf("incorrect status value [false]")
	}

	return items.Items, nil
}