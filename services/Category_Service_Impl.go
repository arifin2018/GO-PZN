package services

import (
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

func (repository *CategoryServiceImpl) Insert() {
}

func (repository *CategoryServiceImpl) Update() {
}

func (repository *CategoryServiceImpl) Delete() {
}
