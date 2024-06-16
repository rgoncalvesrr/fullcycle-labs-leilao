package mongodb

import (
	"context"
	"os"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, logger.Error("erro tentando conectar-se ao MongoDB", err)
	}

	// if err := client.Ping(ctx, nil); err != nil {
	// 	return nil, logger.Error("erro tentando \"pingar\" o MongoDB", err)
	// }

	return client.Database(mongoDatabase), nil
}
