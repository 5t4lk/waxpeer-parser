package csgobackpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BPRequestURL = "http://csgobackpack.net/api/GetItemsList/v2/"
)

func BPRequest() ([]byte, error) {
	r, err := http.Get(BPRequestURL)
	if err != nil {
		return nil, err
	}

	byteInfo, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return byteInfo, nil
}

func BPOrderedOutput(input []byte) (Welcome, error) {
	var r Welcome

	err := json.Unmarshal(input, &r)
	if err != nil {
		fmt.Print("error while unmarshalling...", err)
	}

	return r, err
}
