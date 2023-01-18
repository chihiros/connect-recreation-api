package infra

import (
	"app/ent"
	"app/interfaces/controller"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(conn *ent.Client) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := controller.NewController(conn)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", controller.Get)
			r.Get("/query", controller.GetByID)
			r.Post("/", controller.Post)
			r.Put("/", controller.Put)
			r.Delete("/", controller.Delete)
		})

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
	})

	return r
}
