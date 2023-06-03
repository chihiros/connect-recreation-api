package usecase

import (
	"context"

	"github.com/google/uuid"
)

type RecreationUseCase interface {
	GetRecreations(context.Context, int, int) (Response, error)
	GetRecreationsByID(context.Context, uuid.UUID) (Response, error)
	PostRecreations(context.Context, Request) (Response, error)
	// DeleteRecreationsByID(context.Context, int) error
}

type RecreationRepository interface {
	GetRecreations(context.Context, int, int) (Response, error)
	GetRecreationsByID(context.Context, uuid.UUID) (Response, error)
	PostRecreations(context.Context, Request) (Response, error)
	// DeleteRecreationsByID(context.Context, int) error
}

type RecreationUsecase struct {
	Repository RecreationRepository
}

func (u *RecreationUsecase) GetRecreations(ctx context.Context, limit, offset int) (Response, error) {
	return u.Repository.GetRecreations(ctx, limit, offset)
}

func (u *RecreationUsecase) GetRecreationsByID(ctx context.Context, id uuid.UUID) (Response, error) {
	return u.Repository.GetRecreationsByID(ctx, id)
}

func (u *RecreationUsecase) PostRecreations(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostRecreations(ctx, req)
}

// func (u *RecreationUsecase) DeleteRecreationsByID(ctx context.Context, id int) error {
// 	return u.Repository.DeleteRecreationsByID(ctx, id)
// }
