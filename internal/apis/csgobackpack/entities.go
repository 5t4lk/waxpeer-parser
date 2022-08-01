package csgobackpack

type ItemInfo struct {
	Success           bool   `json:"success"`
	AveragePrice      string `json:"average_price"`
	MedianPrice       string `json:"median_price"`
	AmountSold        string `json:"amount_sold"`
	StandardDeviation string `json:"standard_deviation"`
	LowestPrice       string `json:"lowest_price"`
	HighestPrice      string `json:"highest_price"`
	FirstSaleDate     string `json:"first_sale_date"`
	Time              string `json:"time"`
	Icon              string `json:"icon"`
	Currency          string `json:"currency"`
}
