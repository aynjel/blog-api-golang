package crud_db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *sql.DB
)

func InitDB() {
	var err error
	Client, err = sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic(err)
	}
}
