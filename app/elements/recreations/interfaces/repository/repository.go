package repository

import (
	"app/elements/recreations/usecase"
	"app/ent"
	"app/ent/profile"
	"app/ent/recreation"
	"app/middle/applog"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RecreationRepository struct {
	DBConn *ent.Client
}

func NewRecreationRepository(conn *ent.Client) *RecreationRepository {
	return &RecreationRepository{
		DBConn: conn,
	}
}

type RecreationResponse struct {
	Recreations  []*ent.Recreation `json:"recreations"`
	TotalRecords int               `json:"total_records"`
}

func (r *RecreationRepository) GetRecreations(ctx context.Context, limit, offset int) (usecase.Response, error) {
	// count all records first
	count, err := r.DBConn.Recreation.Query().Count(ctx)
	if err != nil {
		applog.Panic(err)
	}

	// then fetch paged records
	recreation, err := r.DBConn.Recreation.
		Query().
		Order(ent.Desc(recreation.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		applog.Panic(err)
	}

	recRes := RecreationResponse{
		Recreations:  recreation,
		TotalRecords: count,
	}

	res := usecase.Response{
		Data: recRes,
	}
	return res, err
}

func (r *RecreationRepository) GetRecreationsByID(ctx context.Context, id uuid.UUID) (usecase.Response, error) {
	recreation, err := r.DBConn.Recreation.
		Query().
		Where(recreation.RecreationIDEQ(id)).
		Only(ctx)

	if err != nil {
		applog.Panic(err)
	}

	profile, err := r.DBConn.Profile.Query().
		Where(profile.UUIDEQ(recreation.UserID)).
		First(ctx)

	if err != nil {
		applog.Panic(err)
	}

	recreation.Edges.Profile = profile

	res := usecase.Response{Data: recreation}
	return res, err
}

func (r *RecreationRepository) PostRecreations(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	user, err := r.DBConn.Recreation.Create().
		SetUserID(req.UserID).
		SetRecreationID(req.RecreationID).
		SetGenre(req.Genre).
		SetTitle(req.Title).
		SetContent(req.Content).
		SetYoutubeID(req.YouTubeID).
		SetTargetNumber(req.TargetNumber).
		SetRequiredTime(req.RequiredTime).
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

	res := usecase.Response{Data: user}
	return res, err
}

// func (r *RecreationRepository) DeleteRecreationsByID(ctx context.Context, id int) error {
// 	_, err := r.DBConn.Recreation.Delete().
// 		Where(user.IDEQ(id)).
// 		Exec(ctx)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return err
// }
