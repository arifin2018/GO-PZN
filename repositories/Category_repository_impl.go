package repositories

import (
	"GoRestfulApi/helper"
	"GoRestfulApi/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	sqlResult, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)
	id, err := sqlResult.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE `Category` SET Name = ? where Id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	if err != nil {
		// Handle error
		// fmt.Println("Failed to execute query in transaction:", err)
		fmt.Println("Failed to execute query in transaction:", err)
		return category
	}
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE from `Category` where Id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	query := "SELECT * from `Category` where Id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryID)
	helper.PanicIfError(err)
	if rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	}
	return domain.Category{}, errors.New("category not found")
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "select * from category"
	sqlRow, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	var categories []domain.Category
	for sqlRow.Next() {
		category := domain.Category{}
		err := sqlRow.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
