package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/logger"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/user/core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (r *UserRepository) FindUserById(ctx context.Context, userId string) (*entity.User, *internalerror.Error) {

	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo

	err := r.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			msg := fmt.Sprintf("usuário %s não localizado.", userId)
			logger.Error(msg, err)
			return nil, internalerror.NewNotFoundError(msg)
		}

		msg := "erro tentando localizar usuário"
		logger.Error(msg, err)
		return nil, internalerror.NewInternalServerError(msg)
	}

	return &entity.User{
		ID:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}, nil
}
