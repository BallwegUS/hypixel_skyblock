package api

type BazaarData struct {
	Success     bool
	LastUpdated int
	Products    map[string]Product
}

type Product struct {
	ID          string        `json:"product_id"`
	SellSummary []SummaryInfo `json:"sell_summary"`
	BuySummary  []SummaryInfo `json:"buy_summary"`
	QuickStatus QuickStatus   `json:"quick_status"`
}

type SummaryInfo struct {
	Amount       int
	PricePerUnit float64
	Orders       int
}

type QuickStatus struct {
	SellPrice      float64
	SellVolume     int
	SellMovingWeek int
	SellOrders     int
	BuyPrice       float64
	BuyVolume      int
	BuyMovingWeek  int
	BuyOrders      int
}
