package gopzn

import (
	"database/sql"
	"testing"
)

func TestEmpty(t *testing.T)  {
	
}

func TestOpenConnection(t *testing.T)  {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}