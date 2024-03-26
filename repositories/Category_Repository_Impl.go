package repositories

import (
	"context"
)

type CategoryRepositoryImpl struct {
}

func CategoryRepositoryConstruct() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Get(ctx context.Context) {

}

func (repository *CategoryRepositoryImpl) GetById() {

}

func (repository *CategoryRepositoryImpl) Insert() {

}

func (repository *CategoryRepositoryImpl) Update() {

}

func (repository *CategoryRepositoryImpl) Delete() {
}
