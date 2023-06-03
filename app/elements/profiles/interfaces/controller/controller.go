package controller

import (
	"app/elements/profiles/interfaces/repository"
	"app/elements/profiles/usecase"
	"app/ent"
	"app/middle/authrization"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type ProfileController struct {
	Usecase usecase.ProfileUseCase
}

func NewProfileController(conn *ent.Client) *ProfileController {
	u := NewProfileUsecase(conn)
	return &ProfileController{
		Usecase: u,
	}
}

func NewProfileUsecase(conn *ent.Client) *usecase.ProfileUsecase {
	repo := repository.NewProfileRepository(conn)
	return &usecase.ProfileUsecase{
		Repository: repo,
	}
}

func getUUIDWithClaims(r *http.Request) uuid.UUID {
	claims, ok := r.Context().Value("claims").(*authrization.CustomClaims)
	if !ok {
		panic("Invalid user claims")
	}
	return uuid.MustParse(claims.Subject)
}

func (c *ProfileController) GetProfiles(w http.ResponseWriter, r *http.Request) {
	uuid := getUUIDWithClaims(r)
	profiles, err := c.Usecase.GetProfiles(context.Background(), uuid)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(profiles)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profiles)
}

func (c *ProfileController) PostProfiles(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)

	// Request.UUIDを上書きする
	req.UUID = getUUIDWithClaims(r)

	profile, err := c.Usecase.PostProfiles(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(profile)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func (c *ProfileController) PutProfiles(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)

	// Request.UUIDを上書きする
	req.UUID = getUUIDWithClaims(r)

	profile, err := c.Usecase.PutProfiles(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(profile)
		default:
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

func (c *ProfileController) DeleteProfiles(w http.ResponseWriter, r *http.Request) {
	uuid := getUUIDWithClaims(r)
	profile := c.Usecase.DeleteProfiles(context.Background(), uuid)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(profile)
}
