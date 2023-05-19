package main

import (
	"app/infra"
	"app/middle/applog"
	"net/http"
	"time"
)

func main() {
	applog.Setenv(applog.DEV)
	// applog.Setenv(applog.PROD)

	// ロケールを日本に設定する
	{
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			applog.Panic(err)
		}
		time.Local = jst
	}

	// // SQLite3へのコネクションを取得する
	// conn, err := infra.NewSQLite3Connection()
	// if err != nil {
	// 	panic(err)
	// }

	// Postgresへのコネクションを取得する
	conn, err := infra.NewPostgresConnection()
	if err != nil {
		applog.Panic(err)
	}

	r := infra.NewRouter(conn)
	if err := http.ListenAndServe(":8080", r); err != nil {
		applog.Panic(err)
	}
}
