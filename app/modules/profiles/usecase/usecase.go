package usecase

import (
	"context"

	"github.com/google/uuid"
)

type ProfileUseCase interface {
	GetProfiles(context.Context, uuid.UUID) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	PutProfiles(context.Context, Request) (Response, error)
	DeleteProfiles(context.Context, uuid.UUID) error
}

type ProfileRepository interface {
	GetProfiles(context.Context, uuid.UUID) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	PutProfiles(context.Context, Request) (Response, error)
	DeleteProfiles(context.Context, uuid.UUID) error
}

type ProfileUsecase struct {
	Repository ProfileRepository
}

func (u *ProfileUsecase) GetProfiles(ctx context.Context, uuid uuid.UUID) (Response, error) {
	return u.Repository.GetProfiles(ctx, uuid)
}

func (u *ProfileUsecase) PostProfiles(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostProfiles(ctx, req)
}

func (u *ProfileUsecase) PutProfiles(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PutProfiles(ctx, req)
}

func (u *ProfileUsecase) DeleteProfiles(ctx context.Context, uuid uuid.UUID) error {
	return u.Repository.DeleteProfiles(ctx, uuid)
}
