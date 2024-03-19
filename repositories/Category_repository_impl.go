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
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	defer repository.db.Close()
	query := "INSERT INTO `Category` (`Name`) VALUES (?)"
	sqlResult, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)
	id, err := sqlResult.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	defer repository.db.Close()
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
	defer repository.db.Close()
	query := "DELETE from `Category` where Id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	defer repository.db.Close()
	query := "SELECT * from `Category` where Id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryID)
	defer rows.Close()
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
	defer repository.db.Close()
	query := "select * from category"
	sqlRow, err := tx.QueryContext(ctx, query)
	defer sqlRow.Close()
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
