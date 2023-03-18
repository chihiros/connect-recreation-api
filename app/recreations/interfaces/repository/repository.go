package repository

import (
	"app/ent"
	"app/recreations/usecase"
	"context"
	"fmt"
	"time"
)

type RecreationRepository struct {
	DBConn *ent.Client
}

func NewRecreationRepository(conn *ent.Client) *RecreationRepository {
	return &RecreationRepository{
		DBConn: conn,
	}
}

// func (r *RecreationRepository) GetRecreations(ctx context.Context) (usecase.Response, error) {
// 	users, err := r.DBConn.Recreation.Query().All(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	res := usecase.Response{Data: users}
// 	return res, err
// }

// func (r *RecreationRepository) GetRecreationsByID(ctx context.Context, id int) (usecase.Response, error) {
// 	user, err := r.DBConn.Recreation.Query().
// 		Where(user.IDEQ(id)).
// 		All(ctx)

// 	if err != nil {
// 		panic(err)
// 	}

// 	res := usecase.Response{Data: user}
// 	return res, err
// }

func (r *RecreationRepository) PostRecreations(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	user, err := r.DBConn.Recreation.Create().
		SetUID(req.UID).
		SetUsername(req.Username).
		SetMail(req.Mail).
		SetPrefectureID(req.PrefectureID).
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
