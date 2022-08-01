package shadowpay

type Items struct {
	Data     []Item   `json:"data"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
}

type Item struct {
	ID          int           `json:"id"`
	Price       float64       `json:"price"`
	Floatvalue  float64       `json:"floatvalue"`
	Paintindex  int           `json:"paintindex"`
	Paintseed   int           `json:"paintseed"`
	Link        string        `json:"link"`
	TimeCreated string        `json:"time_created"`
	Stickers    []interface{} `json:"stickers"`
	SteamItem   SteamItem     `json:"steam_item"`
}

type SteamItem struct {
	ID                  int         `json:"id"`
	Project             string      `json:"project"`
	SteamMarketHashName string      `json:"steam_market_hash_name"`
	Exterior            string      `json:"exterior"`
	Type                string      `json:"type"`
	Subcategory         string      `json:"subcategory"`
	Collection          string      `json:"collection"`
	Phase               interface{} `json:"phase"`
	SuggestedPrice      float64     `json:"suggested_price"`
	IsStattrak          bool        `json:"is_stattrak"`
	Icon                string      `json:"icon"`
	Rarity              string      `json:"rarity"`
}

type Metadata struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
