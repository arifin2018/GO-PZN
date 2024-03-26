package tests

import (
	"GoRestfulApi/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBTesting() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp/restful_api_test?charset=utf8mb4,utf8")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db
}
