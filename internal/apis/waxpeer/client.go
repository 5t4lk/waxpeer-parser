package waxpeer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	getRequestURL = "https://api.waxpeer.com/v1/prices?game=csgo&minified=steam_price"
)

func GetWXRequest() ([]byte, error) {
	r, err := http.Get(getRequestURL)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}

func GetOrderedOutput(input []byte) ([]ItemInfo, error) {
	var itemsData ItemsData

	if err := json.Unmarshal(input, &itemsData); err != nil {
		return nil, err
	}

	if !itemsData.Success {
		fmt.Println("incorrect success status")
	}

	return itemsData.Items, nil
}
