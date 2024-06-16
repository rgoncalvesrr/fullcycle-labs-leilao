package infra

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/error"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                  `bson:"_id"`
	ProductName string                  `bson:"product_name"`
	Category    string                  `bson:"category"`
	Description string                  `bson:"description"`
	Condition   entity.ProductCondition `bson:"condition"`
	Status      entity.AuctionStatus    `bson:"status"`
	CreatedAt   int64                   `bson:"created_at"`
}

type AuctionRepository struct {
	Collection mongo.Collection
}

func NewAuctionRepository(db *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: *db.Collection("auctions"),
	}
}

func (r *AuctionRepository) Create(ctx context.Context, auction entity.Auction) *error.InternalError {

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		CreatedAt:   auction.CreatedAt.Unix(),
	}

	_, err := r.Collection.InsertOne(ctx, auctionEntityMongo)

	if err != nil {
		msg := "ocorreu um erro tentando inserir um leilão"
		logger.Error(msg, err)
		return error.NewInternalServerError(msg)
	}

	return nil
}
