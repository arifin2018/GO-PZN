package main

import (
	"context"
	"database/sql"
	"fmt"
	databaseconnection "mysql/databaseConnection"
	"mysql/repositories"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection2() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func main()  {
	db := databaseconnection.GetConnection()
	defer db.Close()
	CommentRepository := repositories.CommentRepositoryImpl{
		DB: db,
	}
	ctx := context.Background()
	result,err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for k, v := range result {
		fmt.Println(k,v)
	}
}
