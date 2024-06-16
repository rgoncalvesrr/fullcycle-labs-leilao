package entity

import "time"

type ProductCondition int
type AuctionStatus int

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	CreatedAt   time.Time
}

const (
	Active AuctionStatus = iota + 1
	Completed
)

const (
	New ProductCondition = iota + 1
	Used
	Refurbished
)
