package repository

import (
	"app/elements/users/usecase"
	"app/ent"
	"context"
	"fmt"
	"time"
)

type UserRepository struct {
	DBConn *ent.Client
}

func NewUserRepository(conn *ent.Client) *UserRepository {
	return &UserRepository{
		DBConn: conn,
	}
}

func (r *UserRepository) PostUsers(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	user, err := r.DBConn.User.Create().
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
