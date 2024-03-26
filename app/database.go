package app

import (
	"GoResfulApiPribadi/helpers"
	"database/sql"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp/restful_api?charset=utf8mb4,utf8")
	helpers.PanicIfError(err)
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
