package special_item

import "fmt"

func ItemInfo() (string, string, string, error) {
	var min, max, search string

	enter("Min: ", &min)
	enter("Max: ", &max)
	enter("Name: ", &search)

	return min, max, search, nil
}

func enter(text, scan interface{}) {
	fmt.Print(text)
	fmt.Scan(scan)
}
