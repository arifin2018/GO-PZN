package repositories

import (
	"context"
	"database/sql"
	"errors"
	"mysql/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comments)(entity.Comments, error) {
	script := "INSERT INTO comments(email,comment) VALUES(?,?)"
	statment,err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return comment, err
	}
	result,err := statment.ExecContext(ctx, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id,err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = id
	return comment, nil
}
func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int)(entity.Comments, error) {
	script := "SELECT id,email,comment from comments where id = ? limit 1"
	statment,err := repository.DB.Prepare(script)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}
	rows,err := statment.Query(id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}
	return comment, errors.New("Sorry id "+strconv.Itoa(id)+" doesn't have data")
}
func (repository *CommentRepositoryImpl) FindAll(ctx context.Context)([]entity.Comments, error) {
	var comments []entity.Comments
	script := "select id,email,comment from comments"
	rows,err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := entity.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments , nil
}