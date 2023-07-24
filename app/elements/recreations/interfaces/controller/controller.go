package controller

import (
	"app/elements/recreations/interfaces/repository"
	"app/elements/recreations/usecase"
	"app/ent"
	"app/middle/applog"
	"app/middle/authrization"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		applog.Warn(err.Error())
		limit = 10
	}

	if limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		applog.Warn(err.Error())
		offset = 0
	}

	if offset < 0 {
		offset = 0
	}

	users, err := c.Usecase.GetRecreations(context.Background(), limit, offset)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *RecreationController) GetRecreationsByID(w http.ResponseWriter, r *http.Request) {
	var users usecase.Response
	id := r.URL.Query().Get("id")

	if id == "" {
		c.GetRecreations(w, r)
		return
	}

	// idをUUIDに変換
	recID, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	users, err = c.Usecase.GetRecreationsByID(context.Background(), recID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *RecreationController) PostRecreations(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("%v\n", err)
	}

	// jwtのplayloadからuser_idを取得
	payload, ok := r.Context().Value(authrization.PayloadKey).(*authrization.SupabaseJwtPayload)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user_id, err := uuid.Parse(payload.Subject)
	if err != nil {
		applog.Error(err.Error())
	}

	req.UserID = user_id
	recreation, err := c.Usecase.PostRecreations(context.Background(), req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(recreation)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recreation)
}

func (c *RecreationController) GetRecreationsDraft(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		applog.Warn(err.Error())
		limit = 10
	}

	if limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		applog.Warn(err.Error())
		offset = 0
	}

	if offset < 0 {
		offset = 0
	}

	users, err := c.Usecase.GetRecreationsDraft(
		context.Background(),
		user_id,
		limit,
		offset,
	)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *RecreationController) GetRecreationsDraftByID(w http.ResponseWriter, r *http.Request) {
	// jwtのplayloadからuser_idを取得
	payload, ok := r.Context().Value(authrization.PayloadKey).(*authrization.SupabaseJwtPayload)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user_id, err := uuid.Parse(payload.Subject)
	if err != nil {
		applog.Error(err.Error())
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		c.GetRecreationsDraft(w, r, user_id)
		return
	}

	// idをUUIDに変換
	recID, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	var users usecase.Response
	users, err = c.Usecase.GetRecreationsDraftByID(
		context.Background(),
		recID,
		user_id,
	)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *RecreationController) PutRecreationsDraft(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("%v\n", err)
	}

	// jwtのplayloadからuser_idを取得
	payload, ok := r.Context().Value(authrization.PayloadKey).(*authrization.SupabaseJwtPayload)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	user_id, err := uuid.Parse(payload.Subject)
	if err != nil {
		applog.Error(err.Error())
	}

	req.UserID = user_id
	recreation, err := c.Usecase.PutRecreationsDraft(context.Background(), req)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(recreation)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recreation)
}
