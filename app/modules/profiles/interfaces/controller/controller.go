package controller

import (
	"app/ent"
	"app/middle/applog"
	"app/middle/authrization"
	"app/modules/profiles/interfaces/repository"
	"app/modules/profiles/usecase"
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

func getUUIDWithPayload(r *http.Request) uuid.UUID {
	payload, ok := r.Context().Value(authrization.PayloadKey).(*authrization.SupabaseJwtPayload)
	if !ok {
		applog.Panic(errors.New("Invalid user payload"))
	}
	return uuid.MustParse(payload.Subject)
}

func getPayload(r *http.Request) *authrization.SupabaseJwtPayload {
	payload, ok := r.Context().Value(authrization.PayloadKey).(*authrization.SupabaseJwtPayload)
	if !ok {
		applog.Panic(errors.New("Invalid user payload"))
	}
	return payload
}

func (c *ProfileController) GetProfiles(w http.ResponseWriter, r *http.Request) {
	uuid := getUUIDWithPayload(r)
	profiles, err := c.Usecase.GetProfiles(context.Background(), uuid)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(profiles)
			return
		}

		applog.Panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profiles)
}

func (c *ProfileController) PostProfiles(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)

	// jwtのplayloadからuser_idを取得
	payload := getPayload(r)

	// Request.UUIDを上書きする
	uuid, err := uuid.Parse(payload.Subject)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	req.UUID = uuid
	req.IconURL = payload.UserMetadata.AvatarURL
	if payload.UserMetadata.UserName != "" {
		req.Nickname = payload.UserMetadata.UserName
	} else if payload.UserMetadata.PreferredUsername != "" {
		req.Nickname = payload.UserMetadata.PreferredUsername
	} else if payload.UserMetadata.Name != "" {
		req.Nickname = payload.UserMetadata.Name
	} else if payload.UserMetadata.FullName != "" {
		req.Nickname = payload.UserMetadata.FullName
	} else {
		req.Nickname = ""
	}
	profile, err := c.Usecase.PostProfiles(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(profile)
		default:
			applog.Panic(err)
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
	req.UUID = getUUIDWithPayload(r)

	profile, err := c.Usecase.PutProfiles(context.Background(), req)

	if err != nil {
		switch err.Error() {
		case "duplicate":
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(profile)
		default:
			applog.Panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

func (c *ProfileController) DeleteProfiles(w http.ResponseWriter, r *http.Request) {
	uuid := getUUIDWithPayload(r)
	profile := c.Usecase.DeleteProfiles(context.Background(), uuid)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(profile)
}
