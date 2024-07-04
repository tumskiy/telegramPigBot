package htopd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./store.db")
	if err != nil {
		panic("can not open db")
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id BIGINT, name TEXT);")
	if err != nil {
		return
	}
	fmt.Println(result)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS pd (id BIGINT, user_id BIGINT,chat_id BIGINT, date TIMESTAMP);")
	if err != nil {
		return
	}
	fmt.Println(result)
}
