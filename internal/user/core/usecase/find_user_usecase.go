package usecase

import (
	"context"

	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/internalerror"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/internal/user/core/entity"
)

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type IFindUserUseCase interface {
	Execute(ctx context.Context, userId string) (*UserOutputDTO, *internalerror.Error)
}

type FindUserUseCase struct {
	UserRepository entity.IUserReposiroty
}

func NewFindUserUseCase(userRepository entity.IUserReposiroty) IFindUserUseCase {
	return &FindUserUseCase{
		UserRepository: userRepository,
	}
}

func (u *FindUserUseCase) Execute(ctx context.Context, userId string) (*UserOutputDTO, *internalerror.Error) {
	user, err := u.UserRepository.FindUserById(ctx, userId)

	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}
