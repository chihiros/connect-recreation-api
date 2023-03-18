package repository

import (
	"app/ent"
	"app/ent/user"
	"app/users/usecase"
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

func (r *UserRepository) GetUsers(ctx context.Context) (usecase.Response, error) {
	users, err := r.DBConn.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: users}
	return res, err
}

func (r *UserRepository) GetUsersByID(ctx context.Context, id int) (usecase.Response, error) {
	user, err := r.DBConn.User.Query().
		Where(user.IDEQ(id)).
		All(ctx)

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: user}
	return res, err
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

func (r *UserRepository) DeleteUsersByID(ctx context.Context, id int) error {
	_, err := r.DBConn.User.Delete().
		Where(user.IDEQ(id)).
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	return err
}
