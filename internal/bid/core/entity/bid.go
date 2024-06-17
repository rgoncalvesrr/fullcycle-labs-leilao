package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	CreatedAt time.Time
}

func NewBid(userId, auctionId string, amount float64) (*Bid, *internalerror.Error) {
	bid := &Bid{
		Id:        uuid.NewString(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	if err := bid.Validate(); err != nil {
		return nil, err
	}
	return bid, nil
}

func (b *Bid) Validate() *internalerror.Error {
	if len(b.UserId) == 0 || len(b.AuctionId) == 0 || b.Amount <= 0 {
		return internalerror.NewBadRequestError("Parâmetros inválidos para criação do lance")
	}
	return nil
}
