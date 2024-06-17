package usecase

import (
	"context"
	"time"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type AuctionInputDTO struct {
	ProductName string            `json:"product_name"`
	Category    string            `json:"category"`
	Description string            `json:"description"`
	Condition   ProductCondiction `json:"condiction"`
}

type AuctionOutputDTO struct {
	Id          string            `json:"id"`
	ProductName string            `json:"product_name"`
	Category    string            `json:"category"`
	Description string            `json:"description"`
	Condition   ProductCondiction `json:"condiction"`
	Status      AuctionStatus     `json:"status"`
	CreatedAt   time.Time         `json:"created_at" time_format:"2006-01-02 15:04:05"`
}

type ProductCondiction int64
type AuctionStatus int64

type ICreateAuctionUseCase interface {
	Execute(ctx context.Context, input AuctionInputDTO) (*AuctionOutputDTO, *internalerror.Error)
}

type CreateAuctionUseCase struct {
	auctionRepository entity.IAuctionRepository
}

func NewCreateAuctionUseCase(auctionRepository entity.IAuctionRepository) ICreateAuctionUseCase {
	return &CreateAuctionUseCase{
		auctionRepository: auctionRepository,
	}
}

func (u *CreateAuctionUseCase) Execute(ctx context.Context, input AuctionInputDTO) (*AuctionOutputDTO, *internalerror.Error) {
	auctionEntity, err := entity.NewAuction(
		input.ProductName,
		input.Category,
		input.Description,
		entity.ProductCondition(input.Condition),
	)

	if err != nil {
		return nil, err
	}

	err = u.auctionRepository.Create(ctx, auctionEntity)

	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondiction(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		CreatedAt:   auctionEntity.CreatedAt,
	}, nil
}
