package entity

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type IAuctionRepository interface {
	Create(ctx context.Context, auction *Auction) *internalerror.Error
	FindAuctionById(ctx context.Context, id string) (*Auction, *internalerror.Error)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internalerror.Error)
}
