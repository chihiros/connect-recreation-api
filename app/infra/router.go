package infra

import (
	contact_controller "app/elements/contact/interfaces/controller"
	profile_controller "app/elements/profiles/interfaces/controller"
	rec_controller "app/elements/recreations/interfaces/controller"
	"app/middle/applog"
	"app/middle/authrization"
	"encoding/json"
	"image/color"
	"net/http"
	"strings"
	"time"

	"github.com/chihiros/logger"
	"github.com/fogleman/gg"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"golang.org/x/image/font/opentype"

	_ "embed"
)

//go:embed DINAlternate-Bold.ttf
var font []byte

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.Logger)
	r.Use(middleware.Recoverer)

	// Access-Control-Allow-Originを許可する
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowedOrigins: []string{"http://localhost:3000", "https://stg.topicpost.net"}, // 特定のオリジンのみ許可
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	ccon := contact_controller.NewContactController()
	// Postgresへのコネクションを取得する
	conn1, err := NewPostgresConnection()
	if err != nil {
		applog.Panic(err)
	}
	rcon := rec_controller.NewRecreationController(conn1)

	conn2, err := NewPostgresConnection()
	if err != nil {
		applog.Panic(err)
	}
	pcon := profile_controller.NewProfileController(conn2)
	r.Route("/v1", func(r chi.Router) {
		// プロフィール用のAPI
		r.Route("/profile", func(r chi.Router) {
			r.Use(authrization.AuthMiddleware) // Dockerで開発するときはコメントアウトする

			r.Get("/", pcon.GetProfiles)
			r.Post("/", pcon.PostProfiles)
			r.Put("/", pcon.PutProfiles)
			r.Delete("/", pcon.DeleteProfiles)
		})

		// レクリエーション用のAPI
		r.Route("/recreation", func(r chi.Router) {
			// JWTが不要なやつ
			// 公開されているレクリエーションの一覧を取得するためのAPI
			r.Get("/", rcon.GetRecreationsByID)

			// JWTが必要なやつ
			r.With(authrization.AuthMiddleware).Group(func(r chi.Router) {
				// レクリエーションを投稿するためのAPI
				r.Post("/", rcon.PostRecreations)

				// 下書きのレクリエーションを取得するためのAPI
				r.Get("/draft", rcon.GetRecreationsDraftByID)
				// レクリエーションの途中保存で使うAPI
				r.Put("/draft", rcon.PutRecreationsDraft)
			})
		})

		// お問い合わせ用のAPI
		r.Route("/contact", func(r chi.Router) {
			r.Post("/", ccon.PostContact)
		})

		// OG画像用のAPI
		r.Route("/og", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				// // 1200x630の画像を生成
				dc := gg.NewContext(1200, 630)

				// 左上から右下に向かってグラデーション
				grad := gg.NewLinearGradient(0, 0, 1200, 630)
				grad.AddColorStop(0, color.RGBA{255, 197, 193, 255})
				grad.AddColorStop(0.25, color.RGBA{244, 222, 244, 255})
				grad.AddColorStop(0.6943, color.RGBA{255, 249, 195, 255})
				grad.AddColorStop(1, color.RGBA{206, 249, 255, 255})
				dc.SetFillStyle(grad)
				dc.MoveTo(0, 0)
				dc.LineTo(1200, 0)
				dc.LineTo(1200, 630)
				dc.LineTo(0, 630)
				dc.ClosePath()
				dc.Fill()

				// 図形のサイズと位置を計算
				rectWidth := 1200 - 2*43
				rectHeight := 630 - 2*41
				rectX := 43
				rectY := 41

				// 背景色を設定
				dc.SetColor(color.RGBA{255, 255, 255, 255})
				dc.DrawRoundedRectangle(float64(rectX), float64(rectY), float64(rectWidth), float64(rectHeight), 16)
				dc.Fill()

				// フォントを読み込む
				fontFace, err := opentype.Parse(font)
				if err != nil {
					http.Error(w, "Failed to parse font", http.StatusInternalServerError)
					return
				}

				face, err := opentype.NewFace(fontFace, &opentype.FaceOptions{
					Size: 64,
					DPI:  72,
				})
				if err != nil {
					http.Error(w, "Failed to create font face", http.StatusInternalServerError)
					return
				}
				dc.SetFontFace(face)


				// 画像をレスポンスとして返す
				w.Header().Set("Content-Type", "image/png")
				dc.EncodePNG(w)
			})
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
