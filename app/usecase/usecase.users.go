package usecase

import (
	"context"
)

type UserUseCase interface {
	GetUsers(context.Context) (Response, error)
	GetUsersByID(context.Context, int) (Response, error)
	PostUsers(context.Context, Request) (Response, error)
	DeleteUsersByID(context.Context, int) error
}

type UserRepository interface {
	GetUsers(context.Context) (Response, error)
	GetUsersByID(context.Context, int) (Response, error)
	PostUsers(context.Context, Request) (Response, error)
	DeleteUsersByID(context.Context, int) error
}

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) GetUsers(ctx context.Context) (Response, error) {
	return u.Repository.GetUsers(ctx)
}

func (u *UserUsecase) GetUsersByID(ctx context.Context, id int) (Response, error) {
	return u.Repository.GetUsersByID(ctx, id)
}

func (u *UserUsecase) PostUsers(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostUsers(ctx, req)
}

func (u *UserUsecase) DeleteUsersByID(ctx context.Context, id int) error {
	return u.Repository.DeleteUsersByID(ctx, id)
}
