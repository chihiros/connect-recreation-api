package repository

import (
	"app/elements/profiles/usecase"
	"app/ent"
	"context"
	"fmt"
	"time"
)

type ProfileRepository struct {
	DBConn *ent.Client
}

func NewProfileRepository(conn *ent.Client) *ProfileRepository {
	return &ProfileRepository{
		DBConn: conn,
	}
}

// func (r *ProfileRepository) GetProfilesByID(ctx context.Context, id int) (usecase.Response, error) {
// 	profile, err := r.DBConn.Profile.Query().
// 		Where(profile.IDEQ(id)).
// 		All(ctx)

// 	if err != nil {
// 		panic(err)
// 	}

// 	res := usecase.Response{Data: profile}
// 	return res, err
// }

func (r *ProfileRepository) PostProfiles(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	profile, err := r.DBConn.Profile.Create().
		SetUID(req.UID).
		SetProfilename(req.Profilename).
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

	res := usecase.Response{Data: profile}
	return res, err
}

// func (r *ProfileRepository) DeleteProfilesByID(ctx context.Context, id int) error {
// 	_, err := r.DBConn.Profile.Delete().
// 		Where(profile.IDEQ(id)).
// 		Exec(ctx)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return err
// }
