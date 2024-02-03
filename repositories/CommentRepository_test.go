package repositories

import (
	"context"
	"fmt"
	databaseconnection "mysql/databaseConnection"
	"mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T)  {
	commentRepository := NewCommentRepository(databaseconnection.GetConnection())
	ctx := context.Background()
	comment := entity.Comments{
		Email: "arifi222n@23gmail.com",
		Comment: "turzxc",
	}
	result,err := commentRepository.Insert(ctx,comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T)  {
	commentRepository := NewCommentRepository(databaseconnection.GetConnection())
	ctx := context.Background()
	
	result,err := commentRepository.FindById(ctx,25)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T)  {
	commentRepository := NewCommentRepository(databaseconnection.GetConnection())
	ctx := context.Background()
	
	result,err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for k, v := range result {
		fmt.Println(k,v)
	}
}