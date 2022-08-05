package waxpeer

type ItemsData struct {
	Success bool       `json:"success"`
	Items   []ItemInfo `json:"items"`
}
type ItemInfo struct {
	Name        string `json:"name"`
	Count       int    `json:"count"`
	Min         int    `json:"min"`
	SteamPrice  int    `json:"steam_price"`
	Img         string `json:"img"`
	Type        string `json:"type"`
	RarityColor string `json:"rarity_color"`
}
