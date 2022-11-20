package main

import (
	"app/infra"
	"net/http"
)

func main() {
	// SQLite3へのコネクションを取得する
	conn, err := infra.NewSQLite3Connection()
	if err != nil {
		panic(err)
	}

	r := infra.NewRouter(conn)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
