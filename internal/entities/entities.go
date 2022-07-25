package entities

type ItemsData struct {
	Success bool    `json:"success"`
	Items   []Items `json:"items"`
}

type Items struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Min   int    `json:"min"`
}
