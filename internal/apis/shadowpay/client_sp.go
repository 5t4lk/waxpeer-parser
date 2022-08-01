package shadowpay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.shadowpay.com/api/v2/user/items?token="
)

type Client struct {
	token string
}

func New(token string) *Client {
	return &Client{token: token}
}

func (c *Client) SPRequest() ([]byte, error) {
	r, err := http.Get(URL + c.token)
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

func SPOrderedOutput(input []byte) ([]Item, error) {
	var SPItems Items

	if err := json.Unmarshal(input, &SPItems); err != nil {
		return nil, err
	}

	if SPItems.Status != "success" {
		return nil, fmt.Errorf("incorrect status [not success]")
	}

	return SPItems.Data, nil
}
