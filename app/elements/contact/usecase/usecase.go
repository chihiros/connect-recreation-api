package usecase

import (
	"context"
)

type UserUseCase interface {
	PostUsers(context.Context, Request) (Response, error)
}

type UserRepository interface {
	PostUsers(context.Context, Request) (Response, error)
}

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) PostUsers(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostUsers(ctx, req)
}
