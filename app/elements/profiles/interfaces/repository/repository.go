package repository

import (
	"app/elements/profiles/usecase"
	"app/ent"
	"app/ent/profile"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ProfileRepository struct {
	DBConn *ent.Client
}

func NewProfileRepository(conn *ent.Client) *ProfileRepository {
	return &ProfileRepository{
		DBConn: conn,
	}
}

func (r *ProfileRepository) GetProfiles(ctx context.Context, uuid uuid.UUID) (usecase.Response, error) {
	profile, err := r.DBConn.Profile.Query().
		Where(profile.UUIDEQ(uuid)).
		All(ctx)

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: profile}
	return res, err
}

func (r *ProfileRepository) PostProfiles(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	profile, err := r.DBConn.Profile.Create().
		SetUUID(req.UUID).
		SetNickname(req.Nickname).
		SetIconURL(req.IconURL).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			// ent側の制約エラー
			return usecase.Response{}, fmt.Errorf("duplicate")
		}
	}

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: profile}
	return res, err
}

func (r *ProfileRepository) PutProfiles(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	profile, err := r.DBConn.Profile.Update().
		Where(profile.UUIDEQ(req.UUID)).
		SetNickname(req.Nickname).
		SetIconURL(req.IconURL).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: profile}
	return res, err
}

func (r *ProfileRepository) DeleteProfiles(ctx context.Context, uuid uuid.UUID) error {
	_, err := r.DBConn.Profile.Delete().
		Where(profile.UUIDEQ(uuid)).
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	return err
}
