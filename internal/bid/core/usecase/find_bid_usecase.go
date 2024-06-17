package usecase

import (
	"context"

	auction "github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/usecase"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/bid/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type IFindBidUseCase interface {
	FindByByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internalerror.Error)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internalerror.Error)
}

type FindBidUseCase struct {
	auctionRepository auction.IAuctionRepository
	bidRepository     entity.IBidRepository
}

func NewFindBidUseCase(bidRepository entity.IBidRepository, auctionRepository auction.IAuctionRepository) IFindBidUseCase {
	return &FindBidUseCase{
		bidRepository:     bidRepository,
		auctionRepository: auctionRepository,
	}
}

func (u *FindBidUseCase) FindByByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internalerror.Error) {

	bids, err := u.bidRepository.FindByByAuctionId(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	var result []BidOutputDTO

	for _, bid := range bids {
		result = append(result, BidOutputDTO{
			Id:        bid.Id,
			AuctionId: bid.AuctionId,
			UserId:    bid.UserId,
			Amount:    bid.Amount,
			CreatedAt: bid.CreatedAt,
		})
	}

	return result, nil
}

func (u *FindBidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internalerror.Error) {
	auction, err := u.auctionRepository.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionResult := &usecase.AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   usecase.ProductCondiction(auction.Condition),
		Status:      usecase.AuctionStatus(auction.Status),
		CreatedAt:   auction.CreatedAt,
	}

	bid, err := u.bidRepository.FindWinningBidByAuctionId(ctx, auctionId)

	var bidResult *BidOutputDTO = nil

	if err == nil {
		bidResult = &BidOutputDTO{
			Id:        bid.Id,
			AuctionId: bid.AuctionId,
			UserId:    bid.UserId,
			Amount:    bid.Amount,
			CreatedAt: bid.CreatedAt,
		}
	}

	return &WinningInfoOutputDTO{
		Auction: auctionResult,
		Bid:     bidResult,
	}, nil
}
