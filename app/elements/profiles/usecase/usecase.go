package usecase

import (
	"context"
)

type ProfileUseCase interface {
	GetProfilesByUUID(context.Context, int) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	// DeleteProfilesByID(context.Context, int) error
}

type ProfileRepository interface {
	GetProfilesByUUID(context.Context, int) (Response, error)
	PostProfiles(context.Context, Request) (Response, error)
	// DeleteProfilesByID(context.Context, int) error
}

type ProfileUsecase struct {
	Repository ProfileRepository
}

func (u *ProfileUsecase) GetProfilesByUUID(ctx context.Context, id int) (Response, error) {
	return u.Repository.GetProfilesByUUID(ctx, id)
}

func (u *ProfileUsecase) PostProfiles(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostProfiles(ctx, req)
}

// func (u *ProfileUsecase) DeleteProfilesByID(ctx context.Context, id int) error {
// 	return u.Repository.DeleteProfilesByID(ctx, id)
// }
