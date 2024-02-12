package my_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = "root"
	password = "root"
	host     = "127.0.0.1:3306"
	schema   = "my_db"
	// username = string(os.Getenv("MYSQL_USERNAME"))
	// password = os.Getenv("MYSQL_PASSWORD")
	// host     = os.Getenv("MYSQL_HOST")
	// schema   = os.Getenv("MYSQL_SCHEMA")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
}
