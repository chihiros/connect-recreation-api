package repository

import (
	"app/ent"
	"app/ent/profile"
	"app/ent/recreation"
	"app/middle/applog"
	"app/modules/recreations/usecase"
	"context"
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
	count, err := r.DBConn.Recreation.
		Query().
		Where(recreation.PublishEQ(true)). // 公開されているものだけを取得
		Count(ctx)
	if err != nil {
		applog.Panic(err)
	}

	// then fetch paged records
	recreation, err := r.DBConn.Recreation.
		Query().
		Order(ent.Desc(recreation.FieldCreatedAt)).
		Where(recreation.PublishEQ(true)). // 公開されているものだけを取得
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		applog.Panic(err)
	}

	stack := make(map[uuid.UUID]*ent.Profile)
	for _, rec := range recreation {
		if _, ok := stack[rec.UserID]; !ok {
			profile, err := r.DBConn.Profile.Query().
				Where(profile.UUIDEQ(rec.UserID)).
				First(ctx)

			if err != nil {
				applog.Panic(err)
			}
			stack[rec.UserID] = profile
		}
		rec.Edges.Profile = stack[rec.UserID]
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
		Where(
			recreation.RecreationIDEQ(id),
			recreation.PublishEQ(true), // 公開されていることを確認してから取得
		).
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
	_, err := r.DBConn.Recreation.Create().
		SetUserID(req.UserID).
		SetRecreationID(req.RecreationID).
		SetGenre(req.Genre).
		SetTitle(req.Title).
		SetContent(req.Content).
		SetYoutubeID(req.YouTubeID).
		SetTargetNumber(req.TargetNumber).
		SetRequiredTime(req.RequiredTime).
		SetPublish(true).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetPublishedAt(time.Now()).
		Save(ctx)
		// ↓ 本当はこの用にUpsertを実装したいんだけど、entのバグ？仕様上で[]intの配列が正常に保存できないみたいなので
		// 　Insertをしてみてコンフリクトが起きたらUpdateするようにする
		// OnConflict(
		// 	sql.ConflictColumns(
		// 		recreation.FieldUserID,
		// 		recreation.FieldRecreationID,
		// 	),
		// ).
		// Update(func(ru *ent.RecreationUpsert) {
		// 	ru.UpdateGenre().Set(
		// 		recreation.FieldGenre,
		// 		req.Genre,
		// 	)

		// 	// // r.SetGenre(req.Genre)
		// 	// ru.SetGenre([]int{1, 2, 3})
		// 	ru.SetTitle(req.Title)
		// 	ru.SetContent(req.Content)
		// 	ru.SetYoutubeID(req.YouTubeID)
		// 	ru.SetTargetNumber(req.TargetNumber)
		// 	ru.SetRequiredTime(req.RequiredTime)
		// 	ru.SetPublish(true)
		// 	ru.SetUpdatedAt(time.Now())
		// 	ru.SetPublishedAt(time.Now())
		// }).
		// ID(ctx)

	if err != nil {
		if !ent.IsConstraintError(err) {
			applog.Panic(err)
		}
	}

	if ent.IsConstraintError(err) {
		_, err := r.DBConn.Recreation.Update().
			Where(
				recreation.UserIDEQ(req.UserID),
				recreation.RecreationIDEQ(req.RecreationID),
			).
			SetGenre(req.Genre).
			SetTitle(req.Title).
			SetContent(req.Content).
			SetYoutubeID(req.YouTubeID).
			SetTargetNumber(req.TargetNumber).
			SetRequiredTime(req.RequiredTime).
			SetPublish(true).
			SetUpdatedAt(time.Now()).
			SetPublishedAt(time.Now()).
			Save(ctx)

		if err != nil {
			applog.Panic(err)
		}
	}

	rec, err := r.DBConn.Recreation.Query().
		Where(recreation.RecreationIDEQ(req.RecreationID)).
		Only(ctx)

	if err != nil {
		applog.Panic(err)
	}

	res := usecase.Response{Data: rec}
	return res, err
}

// func (r *RecreationRepository) DeleteRecreationsByID(ctx context.Context, id int) error {
// 	_, err := r.DBConn.Recreation.Delete().
// 		Where(user.IDEQ(id)).
// 		Exec(ctx)

// 	if err != nil {
// 		applog.Panic(err)
// 	}

// 	return err
// }

func (r *RecreationRepository) GetRecreationsDraft(ctx context.Context, user_id uuid.UUID, limit, offset int) (usecase.Response, error) {
	// count all records first
	count, err := r.DBConn.Recreation.
		Query().
		Where(
			recreation.PublishEQ(false),  // 公開されていないものだけを取得
			recreation.UserIDEQ(user_id), // Draftだけは自分のものだけを取得
		).
		Count(ctx)
	if err != nil {
		applog.Panic(err)
	}

	// then fetch paged records
	recreation, err := r.DBConn.Recreation.
		Query().
		Order(ent.Desc(recreation.FieldCreatedAt)).
		Where(
			recreation.PublishEQ(false),  // 公開されていないものだけを取得
			recreation.UserIDEQ(user_id), // Draftだけは自分のものだけを取得
		).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		applog.Panic(err)
	}

	stack := make(map[uuid.UUID]*ent.Profile)
	for _, rec := range recreation {
		if _, ok := stack[rec.UserID]; !ok {
			profile, err := r.DBConn.Profile.Query().
				Where(profile.UUIDEQ(rec.UserID)).
				First(ctx)

			if err != nil {
				applog.Panic(err)
			}
			stack[rec.UserID] = profile
		}
		rec.Edges.Profile = stack[rec.UserID]
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

func (r *RecreationRepository) GetRecreationsDraftByID(ctx context.Context, rec_id, user_id uuid.UUID) (usecase.Response, error) {
	recreation, err := r.DBConn.Recreation.
		Query().
		Where(
			recreation.PublishEQ(false),  // 公開されていないものだけを取得
			recreation.UserIDEQ(user_id), // Draftだけは自分のものだけを取得
			recreation.RecreationIDEQ(rec_id),
		).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return usecase.Response{}, nil
		}

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

func (r *RecreationRepository) PutRecreationsDraft(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	_, err := r.DBConn.Recreation.Create().
		SetUserID(req.UserID).
		SetRecreationID(req.RecreationID).
		SetGenre(req.Genre).
		SetTitle(req.Title).
		SetContent(req.Content).
		SetYoutubeID(req.YouTubeID).
		SetTargetNumber(req.TargetNumber).
		SetRequiredTime(req.RequiredTime).
		SetPublish(false).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
		// OnConflict(
		// 	sql.ConflictColumns(
		// 		recreation.FieldUserID,
		// 		recreation.FieldRecreationID,
		// 	),
		// ).
		// Update(func(r *ent.RecreationUpsert) {
		// 	r.SetGenre(req.Genre)
		// 	r.SetTitle(req.Title)
		// 	r.SetContent(req.Content)
		// 	r.SetYoutubeID(req.YouTubeID)
		// 	r.SetTargetNumber(req.TargetNumber)
		// 	r.SetRequiredTime(req.RequiredTime)
		// 	r.SetPublish(false)
		// 	r.SetUpdatedAt(time.Now())
		// }).
		// ID(ctx)

	if err != nil {
		if !ent.IsConstraintError(err) {
			applog.Panic(err)
		}

		if ent.IsConstraintError(err) {
			_, err := r.DBConn.Recreation.Update().
				Where(
					recreation.UserIDEQ(req.UserID),
					recreation.RecreationIDEQ(req.RecreationID),
				).
				SetGenre(req.Genre).
				SetTitle(req.Title).
				SetContent(req.Content).
				SetYoutubeID(req.YouTubeID).
				SetTargetNumber(req.TargetNumber).
				SetRequiredTime(req.RequiredTime).
				SetPublish(false).
				SetUpdatedAt(time.Now()).
				Save(ctx)

			if err != nil {
				applog.Panic(err)
			}
		}
	}

	rec, err := r.DBConn.Recreation.Query().
		Where(recreation.RecreationIDEQ(req.RecreationID)).
		Only(ctx)

	if err != nil {
		applog.Panic(err)
	}

	res := usecase.Response{Data: rec}
	return res, err
}