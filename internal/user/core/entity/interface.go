package entity

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/error"
)

type IUserReposiroty interface {
	FindUserById(ctx context.Context, userId string) (*User, *error.InternalError)
}
