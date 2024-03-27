package services

import (
	"GoResfulApiPribadi/helpers"
	"GoResfulApiPribadi/model"
	"GoResfulApiPribadi/repositories"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
}

func CategoryServiceConstruct(repository repositories.CategoryRepository, db *sql.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
	}
}

func (repository *CategoryServiceImpl) Get(ctx context.Context) {
	sqltx, _ := repository.DB.Begin()
	repositories.CategoryRepositoryConstruct().Get(ctx, sqltx)
}

func (repository *CategoryServiceImpl) GetById() {
}

func (repository *CategoryServiceImpl) Insert(ctx context.Context, category *model.Category) *model.Category {
	sqltx, err := repository.DB.Begin()
	helpers.PanicIfError(err)

	defer func() {
		if err != nil {
			errorRollback := sqltx.Rollback()
			helpers.PanicIfError(errorRollback)
		} else {
			errorCommit := sqltx.Commit()
			helpers.PanicIfError(errorCommit)
		}
	}()

	dataResultCategory := repositories.CategoryRepositoryConstruct().Insert(ctx, sqltx, category)
	return dataResultCategory
}

func (repository *CategoryServiceImpl) Update() {
}

func (repository *CategoryServiceImpl) Delete() {
}
