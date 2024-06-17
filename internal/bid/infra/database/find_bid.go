package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/bid/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *BidRepository) FindByByAuctionId(ctx context.Context, auctionId string) ([]entity.Bid, *internalerror.Error) {

	filter := bson.M{"auction_id": auctionId}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		msg := fmt.Sprintf("Erro ao recuperar lances do leilão %s", auctionId)
		logger.Error(msg, err)

		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, internalerror.NewNotFoundError(msg)
		}

		return nil, internalerror.NewInternalServerError(msg)
	}

	var bidsMongo []BidEntityMongo

	if err := cursor.All(ctx, &bidsMongo); err != nil {
		return nil, internalerror.NewInternalServerError(err.Error())
	}

	bids := []entity.Bid{}

	for _, bid := range bidsMongo {
		bids = append(bids, entity.Bid{
			Id:        bid.Id,
			AuctionId: bid.AuctionId,
			UserId:    bid.UserId,
			Amount:    bid.Amount,
			CreatedAt: time.Unix(bid.CreatedAt, 0),
		})
	}

	return bids, nil
}

func (r *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*entity.Bid, *internalerror.Error) {
	filter := bson.M{"auction_id": auctionId}
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})

	var bidMongo BidEntityMongo

	err := r.Collection.FindOne(ctx, filter, opts).Decode(&bidMongo)
	if err != nil {
		msg := fmt.Sprintf("Erro ao recuperar maior lance do leilão %s", auctionId)
		logger.Error(msg, err)
		return nil, internalerror.NewInternalServerError(msg)
	}

	return &entity.Bid{
		Id:        bidMongo.Id,
		AuctionId: bidMongo.AuctionId,
		UserId:    bidMongo.UserId,
		Amount:    bidMongo.Amount,
		CreatedAt: time.Unix(bidMongo.CreatedAt, 0),
	}, nil
}
