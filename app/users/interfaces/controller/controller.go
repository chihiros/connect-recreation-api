package controller

import (
	"app/ent"
	"app/users/interfaces/repository"
	"app/users/usecase"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Usecase.GetUsers(context.Background())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUsersByID(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからidを取得する
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	users, err := c.Usecase.GetUsersByID(context.Background(), id)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
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

func (c *UserController) DeleteUsersByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	user := c.Usecase.DeleteUsersByID(context.Background(), id)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(user)
}
