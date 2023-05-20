package controller

import (
	"app/elements/recreations/interfaces/repository"
	"app/elements/recreations/usecase"
	"app/ent"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RecreationController struct {
	Usecase usecase.RecreationUseCase
}

func NewRecreationController(conn *ent.Client) *RecreationController {
	u := NewRecreationUsecase(conn)
	return &RecreationController{
		Usecase: u,
	}
}

func NewRecreationUsecase(conn *ent.Client) *usecase.RecreationUsecase {
	repo := repository.NewRecreationRepository(conn)
	return &usecase.RecreationUsecase{
		Repository: repo,
	}
}

func (c *RecreationController) GetRecreations(w http.ResponseWriter, r *http.Request) {
	users, err := c.Usecase.GetRecreations(context.Background())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// func (c *RecreationController) GetRecreationsByID(w http.ResponseWriter, r *http.Request) {
// 	// クエリパラメータからidを取得する
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

// 	users, err := c.Usecase.GetRecreationsByID(context.Background(), id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(users)
// }

func (c *RecreationController) PostRecreations(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("%v\n", err)
	}

	user, err := c.Usecase.PostRecreations(context.Background(), req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

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

// func (c *RecreationController) DeleteRecreationsByID(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
// 	user := c.Usecase.DeleteRecreationsByID(context.Background(), id)

// 	w.WriteHeader(http.StatusNoContent)
// 	json.NewEncoder(w).Encode(user)
// }
