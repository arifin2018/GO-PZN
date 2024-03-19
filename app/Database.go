package app

import (
	"GoRestfulApi/helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp/restful_api?charset=utf8mb4,utf8")
	helper.PanicIfError(err)

	return db
}
