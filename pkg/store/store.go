package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:123456@tcp(192.168.220.129:3306)/pps")
	if err != nil {

		panic(err)
	} else {
		DB = conn

	}

}
