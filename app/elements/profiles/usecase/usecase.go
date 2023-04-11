package usecase

import (
	"context"

	"github.com/google/uuid"
)

type ProfileUseCase interface {
	GetProfilesByUUID(context.Context, uuid.UUID) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	DeleteProfilesByUUID(context.Context, uuid.UUID) error
}

type ProfileRepository interface {
	GetProfilesByUUID(context.Context, uuid.UUID) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	DeleteProfilesByUUID(context.Context, uuid.UUID) error
}

type ProfileUsecase struct {
	Repository ProfileRepository
}

func (u *ProfileUsecase) GetProfilesByUUID(ctx context.Context, uuid uuid.UUID) (Response, error) {
	return u.Repository.GetProfilesByUUID(ctx, uuid)
}

func (u *ProfileUsecase) PostProfiles(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostProfiles(ctx, req)
}

func (u *ProfileUsecase) DeleteProfilesByUUID(ctx context.Context, uuid uuid.UUID) error {
	return u.Repository.DeleteProfilesByUUID(ctx, uuid)
}
