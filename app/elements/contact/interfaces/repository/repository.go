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
		if ent.IsConstraintError(err) {
			// ent側の制約エラー
			return usecase.Response{}, fmt.Errorf("duplicate")
		}
	}

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: "success"}
	return res, err
}
