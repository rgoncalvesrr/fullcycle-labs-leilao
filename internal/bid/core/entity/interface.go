package entity

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type IBidRepository interface {
	Create(ctx context.Context, bids []Bid) *internalerror.Error
	FindByByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internalerror.Error)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internalerror.Error)
}
