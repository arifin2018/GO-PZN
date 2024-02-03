package databaseconnection

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecContextInsertSql(t *testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id,email) VALUES(?,?)"
	// _,db := db.Exec(ctx, script)
	_,err := db.ExecContext(ctx, script,"arifin","arifin")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new Customer")
}

func TestQueryContextSelectSql(t *testing.T)  {
	var id,email string
	var balance sql.NullInt64
	var rating sql.NullFloat64
	var created_at,birth_date sql.NullTime
	var married bool


	db := GetConnection()
	defer db.Close()
	// ctx := context.Background()

	script := "SELECT * FROM customer";
	rows, err := db.Query(script)
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		if err := rows.Scan(&id,&email,&balance,&rating,&created_at,&birth_date,&married); err != nil{
			panic(err)
		}
		if balance.Valid && rating.Valid && created_at.Valid && birth_date.Valid {
			fmt.Println(id,email,balance.Int64,rating.Float64,created_at.Time,birth_date.Time,married)
		}
	}
}

func TestQueryContextSelectSqlInjection(t *testing.T)  {
	

	db := GetConnection()
	defer db.Close()
	
	username := "arifin"
	password := "password"

	// script := "SELECT * FROM user WHERE username = '" +username+ "' AND password = '" + password +"' LIMIT 1";
	script := "SELECT * FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.Query(script,username,password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next(){
		var username,password string
		if err := rows.Scan(&username,&password); err != nil{
			panic(err)
		}
		fmt.Println(username)
		fmt.Println("username ditemukan")
	}else{
		fmt.Println("username tidak ditemukan")
	}
}

func TestAutoIncrement(t *testing.T)  {
	db := GetConnection()
	defer db.Close()
	
	email := "arifin@gmail.com"
	comment := "password"

	// script := "SELECT * FROM user WHERE username = '" +username+ "' AND password = '" + password +"' LIMIT 1";
	script := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := db.Exec(script,email,comment)
	if err != nil {
		panic(err)
	}
	resultId,err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success insert new comment with id", resultId)
}

func TestPrepareStatement(t *testing.T)  {
	db := GetConnection()
	defer db.Close()

	script := "INSERT INTO comments(email,comment) VALUES(?,?)"
	statement, err := db.Prepare(script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("arifin%v@gmail.com",i)
		comment := fmt.Sprintf("ini comment ke-%v",i)
		result, err := statement.Exec(email, comment)
		if err != nil {
			panic(err)
		}
		if id, err := result.LastInsertId(); err != nil {
			panic(err)
		}else{
			fmt.Println(id)
		}
	}
}

func TestTransactionDB(t *testing.T)  {
	db := GetConnection()
	defer db.Close();

	ctx,cancel := context.WithTimeout(context.Background(), time.Second * 6)
	defer cancel()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comments(email,comment) VALUES(?,?)"
	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("arifin%v@gmail.com",i)
		comment := fmt.Sprintf("ini comment ke-%v",i)
		result, err := tx.ExecContext(ctx,script,email, comment)

		if err != nil {
			tx.Rollback()
			panic(err)
		}
		if id, err := result.LastInsertId(); err != nil {
			tx.Rollback()
			panic(err)
		}else{
			fmt.Println(id)
		}
	}

	if err := tx.Commit(); err !=nil {
		tx.Rollback()
		panic(err)
	}
}