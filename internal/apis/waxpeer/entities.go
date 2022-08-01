package waxpeer

type Data struct {
	Success bool   `json:"success"`
	Items   []Item `json:"items"`
}

type Item struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Min   int    `json:"min"`
}
