package repository

import (
	"app/elements/contact/usecase"
	"app/ent"
	"context"
	"errors"
	"fmt"
)

type ContactRepository struct{}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func (c *ContactRepository) PostContact(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	err := errors.New("error")

	if err != nil {
		applog.Panic(err)
	}

	res := usecase.Response{Data: "success"}
	return res, err
}
