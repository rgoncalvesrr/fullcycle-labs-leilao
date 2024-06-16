package entity

import "time"

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	CreatedAt time.Time
}
