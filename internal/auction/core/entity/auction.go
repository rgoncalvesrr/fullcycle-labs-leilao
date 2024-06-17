package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

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

func NewAuction(productName, category, description string, condition ProductCondition) (*Auction, *internalerror.Error) {
	auction := &Auction{
		Id:          uuid.NewString(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		CreatedAt:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() *internalerror.Error {
	if len(a.Description) <= 1 || len(a.Category) <= 1 ||
		len(a.Description) <= 10 || !(a.Condition >= New && a.Condition <= Refurbished) {
		return internalerror.NewBadRequestError("Leilão com dados inválidos")
	}

	return nil
}
