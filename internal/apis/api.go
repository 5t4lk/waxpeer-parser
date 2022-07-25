package apis

import (
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
