package repository

import (
	"app/elements/profiles/usecase"
	"app/ent"
	"app/ent/profile"
	"app/middle/applog"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
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
	p, err := r.DBConn.Profile.Query().
		Where(profile.UUIDEQ(uuid)).
		Only(ctx)

	if err != nil {
		// NotFoundだったら
		if ent.IsNotFound(err) {
			return usecase.Response{
					Data: nil,
					ErrorResponse: usecase.ErrorResponse{
						ErrorCode:    "404",
						ErrorMessage: "not found",
					},
				},
				fmt.Errorf("not found")
		}
	}

	if err != nil {
		applog.Panic(err)
	}

	res := usecase.Response{Data: p}
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
	_, err := r.DBConn.Profile.Create().
		SetUUID(req.UUID).
		SetNickname(req.Nickname).
		SetIconURL(req.IconURL).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		OnConflict(
			sql.ConflictColumns(profile.FieldUUID),
		).
		Update(func(p *ent.ProfileUpsert) {
			p.SetNickname(req.Nickname)
			p.SetIconURL(req.IconURL)
			p.SetUpdatedAt(time.Now())
		}).
		ID(ctx)

	if err != nil {
		panic(err)
	}

	// 更新に成功したら、更新後のデータを返す
	profile, err := r.DBConn.Profile.Query().
		Where(profile.UUIDEQ(req.UUID)).
		Only(ctx)

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
