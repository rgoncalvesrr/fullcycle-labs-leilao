package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*entity.Auction, *error.InternalError) {
	filter := bson.M{"_id": id}
	var auctionMongo AuctionEntityMongo

	if err := r.Collection.FindOne(ctx, filter).Decode(&auctionMongo); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			msg := fmt.Sprintf("leilão \"%s\" não localizado no banco de dados", id)
			logger.Error(msg, err)
			return nil, error.NewNotFoundError(msg)
		}

		msg := "erro tentando localizar leilão"
		logger.Error(msg, err)
		return nil, error.NewInternalServerError(msg)
	}

	return &entity.Auction{
		Id:          auctionMongo.Id,
		ProductName: auctionMongo.ProductName,
		Category:    auctionMongo.Category,
		Description: auctionMongo.Description,
		Condition:   auctionMongo.Condition,
		Status:      auctionMongo.Status,
		CreatedAt:   time.Unix(auctionMongo.CreatedAt, 0),
	}, nil

}

func (r *AuctionRepository) FindAuctions(
	ctx context.Context,
	status entity.AuctionStatus,
	category, productName string) ([]entity.Auction, *error.InternalError) {
	filter := bson.M{}
	if status >= entity.Active && status <= entity.Completed {
		filter["status"] = status
	}
	if category != "" {
		filter["category"] = category
	}
	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		msg := "erro tentando localizar leilões"
		logger.Error(msg, err)
		return nil, error.NewInternalServerError(msg)
	}
	defer cursor.Close(ctx)

	var auctions []AuctionEntityMongo

	if err := cursor.All(ctx, auctions); err != nil {
		msg := "erro tentando localizar leilões"
		logger.Error(msg, err)
		return nil, error.NewInternalServerError(msg)
	}

	result := []entity.Auction{}

	for _, auctionEntityMongo := range auctions {
		result = append(result, entity.Auction{
			Id:          auctionEntityMongo.Id,
			Category:    auctionEntityMongo.Category,
			ProductName: auctionEntityMongo.ProductName,
			Description: auctionEntityMongo.Description,
			Condition:   auctionEntityMongo.Condition,
			Status:      auctionEntityMongo.Status,
			CreatedAt:   time.Unix(auctionEntityMongo.CreatedAt, 0),
		})
	}

	return result, nil
}
