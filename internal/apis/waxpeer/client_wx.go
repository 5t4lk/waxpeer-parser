package waxpeer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	requestURL = "https://api.waxpeer.com/v1/prices?game=csgo&min_price=1000&max_price=100000000&search="
)

func Request() ([]byte, error) {
	r, err := http.Get(requestURL)
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

func OrderedOutput(input []byte, item string) ([]Item, error) {
	var data Data

	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	if !data.Success {
		return nil, fmt.Errorf("incorrect status value [false]")
	}

	return data.Items, nil
}
