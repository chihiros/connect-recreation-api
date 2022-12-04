package main

import (
	"app/infra"
	"fmt"
	"net/http"
)

func main() {
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

	sql := "UPDATE users SET name = $1 WHERE id = $2;"
	fmt.Printf("sql: %s", sql)

	r := infra.NewRouter(conn)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
