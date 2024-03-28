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

	defer helpers.CommitOrRollback(sqltx, err)

	dataResultCategory := repositories.CategoryRepositoryConstruct().Insert(ctx, sqltx, category)
	return dataResultCategory
}

func (repository *CategoryServiceImpl) Update(ctx context.Context, category *model.Category) *model.Category {
	sqlTx, err := repository.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(sqlTx, err)

	modelCategory := repository.CategoryRepository.Update(ctx, sqlTx, category)
	return modelCategory
}

func (repository *CategoryServiceImpl) Delete() {
}
