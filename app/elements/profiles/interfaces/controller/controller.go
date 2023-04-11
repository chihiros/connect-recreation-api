package controller

import (
	"app/elements/profiles/interfaces/repository"
	"app/elements/profiles/usecase"
	"app/ent"
	"context"
	"encoding/json"
	"net/http"
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

// func (c *ProfileController) GetProfilesByID(w http.ResponseWriter, r *http.Request) {
// 	// クエリパラメータからidを取得する
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

// 	profiles, err := c.Usecase.GetProfilesByID(context.Background(), id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(profiles)
// }

func (c *ProfileController) PostProfiles(w http.ResponseWriter, r *http.Request) {
	// bodyの中身をbindする
	req := usecase.Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
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

// func (c *ProfileController) DeleteProfilesByID(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
// 	profile := c.Usecase.DeleteProfilesByID(context.Background(), id)

// 	w.WriteHeader(http.StatusNoContent)
// 	json.NewEncoder(w).Encode(profile)
// }
