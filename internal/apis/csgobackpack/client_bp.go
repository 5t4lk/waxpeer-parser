package csgobackpack

import (
	"io/ioutil"
	"net/http"
)

const (
	requestURL = "http://csgobackpack.net/api/GetItemPrice/?currency=USD&id="
	endOfURL   = "&time=7&icon=1"
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
