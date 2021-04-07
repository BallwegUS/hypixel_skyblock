package api

type AuctionReturn struct {
	Success       bool
	Page          int
	TotalPages    int
	TotalAuctions int
	LastUpdated   int
	Auctions      []AuctionData
}

type AuctionData struct {
	ID          string `json:"uuid";gorm:"primary_key"`
	Auctioneer  string
	ProfileID   string `json:"profile_id"`
	Start       int
	End         int
	Name        string `json:"item_name"`
	Lore        string `json:"item_lore"`
	Extra       string
	Category    string
	Tier        string
	StartingBid int `json:"starting_bid"`
	Claimed     bool
	HighestBid  int  `json:"highest_bid_amount"`
	BIN         bool `json:"bin"`
	Bids        []BidData
}

type BidData struct {
	AuctionID string `json:"auction_id"`
	Bidder    string
	ProfileID string `json:"profile_id"`
	Amount    int
	Timestamp int
}
