package client

import (
	"444/internal/apis"
	"444/internal/entities"
	"444/internal/special_item"
	"encoding/json"
	"fmt"
)

func Get() ([]byte, error) {
	minPrice, maxPrice, itemName, err := special_item.ItemInfo()
	if err != nil {
		return nil, err
	}

	byteJson, err := apis.Request(minPrice, maxPrice, itemName)
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
