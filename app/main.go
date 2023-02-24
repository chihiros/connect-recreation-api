package main

import (
	"app/infra"
	"app/middle/applog"
	"net/http"
	"time"
)

func main() {
	applog.SetEnv("staging")
	alog := applog.NewLog()

	// // SQLite3へのコネクションを取得する
	// conn, err := infra.NewSQLite3Connection()
	// if err != nil {
	// 	panic(err)
	// }

	// Postgresへのコネクションを取得する
	conn, err := infra.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	r := infra.NewRouter(conn)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}
