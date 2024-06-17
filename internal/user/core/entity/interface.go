package entity

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
)

type IUserReposiroty interface {
	FindUserById(ctx context.Context, userId string) (*User, *internalerror.Error)
}
