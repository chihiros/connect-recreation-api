package usecase

import (
	"context"
)

type ContactUseCase interface {
	PostContact(context.Context, Request) (Response, error)
}

type ContactRepository interface {
	PostContact(context.Context, Request) (Response, error)
}

type ContactUsecase struct {
	Repository ContactRepository
}

func (u *ContactUsecase) PostContact(ctx context.Context, req Request) (Response, error) {
	return u.Repository.PostContact(ctx, req)
}
