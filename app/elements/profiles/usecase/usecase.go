package usecase

import (
	"context"
)

type ProfileUseCase interface {
	GetProfiles(context.Context) (Response, error)
	GetProfilesByID(context.Context, int) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	DeleteProfilesByID(context.Context, int) error
}

type ProfileRepository interface {
	GetProfiles(context.Context) (Response, error)
	GetProfilesByID(context.Context, int) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	DeleteProfilesByID(context.Context, int) error
}

type ProfileUsecase struct {
	Repository ProfileRepository
}

func (u *ProfileUsecase) GetProfiles(ctx context.Context) (Response, error) {
	return u.Repository.GetProfiles(ctx)
}

func (u *ProfileUsecase) GetProfilesByID(ctx context.Context, id int) (Response, error) {
	return u.Repository.GetProfilesByID(ctx, id)
}

func (u *ProfileUsecase) PostProfiles(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostProfiles(ctx, req)
}

func (u *ProfileUsecase) DeleteProfilesByID(ctx context.Context, id int) error {
	return u.Repository.DeleteProfilesByID(ctx, id)
}
