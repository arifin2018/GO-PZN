package services

import (
	"GoResfulApiPribadi/repositories"
	"context"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
}

func CategoryServiceConstruct(repository repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
	}
}

func (repository *CategoryServiceImpl) Get(ctx context.Context) {
	repositories.CategoryRepositoryConstruct().Get(ctx)
}

func (repository *CategoryServiceImpl) GetById() {
}

func (repository *CategoryServiceImpl) Insert() {
}

func (repository *CategoryServiceImpl) Update() {
}

func (repository *CategoryServiceImpl) Delete() {
}
