package api

import "time"

type AuctionReturn struct {
	Success       bool
	Page          int
	TotalPages    int
	TotalAuctions int
	LastUpdated   int
	Auctions      []AuctionData
}

type AuctionCache struct {
	LastUpdated time.Time
	Auctions    []AuctionData
}

type AuctionData struct {
	ID             string `json:"uuid" gorm:"primary_key"`
	Auctioneer     string
	ProfileID      string `json:"profile_id"`
	Start          int
	End            int
	Name           string `json:"item_name"`
	Lore           string `json:"item_lore"`
	Extra          string
	Category       string
	Tier           string
	StartingBid    int `json:"starting_bid"`
	Claimed        bool
	HighestBid     int       `json:"highest_bid_amount" gorm:"index"`
	BIN            bool      `json:"bin"`
	Bids           []BidData `gorm:"foreignKey:AuctionID"`
	FinalSalePrice int
}

type BidData struct {
	AuctionID string `json:"auction_id" gorm:"primary_key;autoIncrement:false"`
	Bidder    string
	ProfileID string `json:"profile_id"`
	Amount    int
	Timestamp int `gorm:"primary_key;autoIncrement:false"`
}

type EndedAuctionReturn struct {
	Success     bool
	LastUpdated int
	Auctions    []EndedAuction
}

type EndedAuction struct {
	AuctionID     string `json:"auction_id"`
	Seller        string
	SellerProfile string
	Buyer         string
	Timestamp     int
	Price         int
	BIN           bool `json:"bin"`
}

func (auction AuctionData) GetHighestBid() BidData {
	highest := BidData{Amount: 0}
	for _, bid := range auction.Bids {
		if bid.Amount > highest.Amount {
			highest = bid
		}
	}
	return highest
}
