package gopzn

import (
	"database/sql"
	"testing"
	"time"
)

func GetConnection(t *testing.T)  {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
}