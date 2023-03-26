package controller

import (
	"app/elements/contact/interfaces/repository"
	"app/elements/contact/usecase"
	"context"
	"encoding/json"
	"net/http"
)

type ContactController struct {
	Usecase usecase.ContactUseCase
}

func NewContactController() *ContactController {
	u := NewContactUsecase()
	return &ContactController{
		Usecase: u,
	}
}

func NewContactUsecase() *usecase.ContactUsecase {
	repo := repository.NewContactRepository()
	return &usecase.ContactUsecase{
		Repository: repo,
	}
}

func (c *ContactController) PostContact(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	user, err := c.Usecase.PostContact(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(user)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
