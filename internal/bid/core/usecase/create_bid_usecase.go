package usecase

import (
	"context"
	"time"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/usecase"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/bid/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type WinningInfoOutputDTO struct {
	Auction *usecase.AuctionOutputDTO `json:"auction"`
	Bid     *BidOutputDTO             `json:"bid,omitempty"`
}

type BidInputDTO struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"create_at" time_format:"2006-01-02 15:04:05"`
}

type ICreateBidUseCase interface {
	Create(ctx context.Context, bids []BidInputDTO) *internalerror.Error
}

type CreateBidUseCase struct {
	bidRepository entity.IBidRepository
}

func NewCreateBidUseCase(bidRepository entity.IBidRepository) IFindBidUseCase {
	return &FindBidUseCase{
		bidRepository: bidRepository,
	}
}

func (u *CreateBidUseCase) Create(ctx context.Context, bids []BidInputDTO) *internalerror.Error {
	// u.bidRepository.Create(ctx)

	return nil
}
