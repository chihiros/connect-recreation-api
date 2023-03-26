package repository

import (
	"app/elements/contact/usecase"
	"app/middle/applog"
	"context"
	"os"
)

type ContactRepository struct{}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{}
}

func (c *ContactRepository) PostContact(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	webhook_url := os.Getenv("WEBHOOK_URL")

	if err != nil {
		applog.Panic(err)
	}

	res := usecase.Response{Data: "success"}
	return res, err
}
