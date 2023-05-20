package infra

import (
	contact_controller "app/elements/contact/interfaces/controller"
	profile_controller "app/elements/profiles/interfaces/controller"
	rec_controller "app/elements/recreations/interfaces/controller"
	user_controller "app/elements/users/interfaces/controller"
	"app/ent"
	"app/middle/authrization"
	"encoding/json"
	"net/http"
	"time"

	"github.com/chihiros/logger"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(conn *ent.Client) *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.Logger)
	r.Use(middleware.Recoverer)

	// Access-Control-Allow-Originを許可する
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	ucon := user_controller.NewUserController(conn)
	ccon := contact_controller.NewContactController()
	rcon := rec_controller.NewRecreationController(conn)
	pcon := profile_controller.NewProfileController(conn)
	r.Route("/v1", func(r chi.Router) {
		// お問い合わせ用のAPI
		r.Route("/contact", func(r chi.Router) {
			r.Post("/", ccon.PostContact)
		})

		// プロフィール用のAPI
		r.Route("/profile", func(r chi.Router) {
			r.Use(authrization.AuthMiddleware) // Dockerで開発するときはコメントアウトする

			r.Get("/", pcon.GetProfiles)
			r.Post("/", pcon.PostProfiles)
			r.Put("/", pcon.PutProfiles)
			r.Delete("/", pcon.DeleteProfiles)
		})

		// ダミーで使っていたAPI（いずれ削除されると思う）
		r.Route("/users", func(r chi.Router) {
			r.Get("/", ucon.GetUsers)
			r.Get("/query", ucon.GetUsersByID)
			r.Post("/", ucon.PostUsers)
			r.Delete("/", ucon.DeleteUsersByID)
		})

		// 疎通確認用のAPI
		r.Route("/now", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				jst, err := time.LoadLocation("Asia/Tokyo")
				if err != nil {
					panic(err)
				}
				now := time.Now().In(jst)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(now)
			})
		})

		// レクリエーション用のAPI
		r.Route("/recreation", func(r chi.Router) {
			// JWTが不要なやつ
			r.Get("/", rcon.GetRecreations)
			r.Get("/query", ucon.GetUsersByID)

			r.Post("/", rcon.PostRecreations)
			// JWTが必要なやつ
			r.With(authrization.AuthMiddleware).Group(func(r chi.Router) {
				r.Delete("/", ucon.DeleteUsersByID)
			})
		})

		// Example
		r.Route("/example", func(r chi.Router) {
			r.Get("/nojwt", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode("This is NOT JWT protected API.")
			})

			r.With(authrization.AuthMiddleware).Group(func(r chi.Router) {
				r.Get("/jwt", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode("This is JWT protected API.")
				})
			})
		})
	})

	return r
}
