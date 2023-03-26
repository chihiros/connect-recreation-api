package controller

import (
	"app/elements/users/interfaces/repository"
	"app/elements/users/usecase"
	"app/ent"
	"context"
	"encoding/json"
	"net/http"
)

type UserController struct {
	Usecase usecase.UserUseCase
}

func NewUserController(conn *ent.Client) *UserController {
	u := NewUserUsecase(conn)
	return &UserController{
		Usecase: u,
	}
}

func NewUserUsecase(conn *ent.Client) *usecase.UserUsecase {
	repo := repository.NewUserRepository(conn)
	return &usecase.UserUsecase{
		Repository: repo,
	}
}

func (c *UserController) PostUsers(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	user, err := c.Usecase.PostUsers(context.Background(), req)

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
