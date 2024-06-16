package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	auction_entity "github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/auction/infra/database"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/bid/core/entity"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	CreatedAt int64   `bson:"created_at"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *database.AuctionRepository
}

func (r *BidRepository) Create(ctx context.Context, bids []entity.Bid) *error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bids {
		wg.Add(1)

		go func(value *entity.Bid) {
			defer wg.Done()
			auction, err := r.AuctionRepository.FindAuctionById(ctx, value.AuctionId)
			if err != nil {
				logger.Error(err.Message, err)
				return
			}
			if auction.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				Id:        value.Id,
				AuctionId: value.AuctionId,
				UserId:    value.UserId,
				Amount:    value.Amount,
				CreatedAt: value.CreatedAt.Unix(),
			}

			if _, err := r.Collection.InsertOne(ctx, &bidEntityMongo); err != nil {
				logger.Error(fmt.Sprintf("tentativa de registro do lance %s falhou", value.Id), err)
				return
			}
		}(&bid)
	}

	wg.Wait()
	return nil
}
