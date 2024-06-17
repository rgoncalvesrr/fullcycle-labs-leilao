package usecase

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type IFindAuctionUseCase interface {
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internalerror.Error)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internalerror.Error)
}

type FindAuctionUseCase struct {
	auctionRepository entity.IAuctionRepository
}

func NewFindAuctionUseCase(auctionRepository entity.IAuctionRepository) IFindAuctionUseCase {
	return &FindAuctionUseCase{
		auctionRepository: auctionRepository,
	}
}

func (u *FindAuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internalerror.Error) {
	auction, err := u.auctionRepository.FindAuctionById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondiction(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		CreatedAt:   auction.CreatedAt,
	}, nil
}

func (u *FindAuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internalerror.Error) {
	auctions, err := u.auctionRepository.FindAuctions(ctx, entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}
	var result []AuctionOutputDTO

	for _, auction := range auctions {
		result = append(result, AuctionOutputDTO{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductCondiction(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			CreatedAt:   auction.CreatedAt,
		})
	}

	return result, nil
}
