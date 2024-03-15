package repositories

import (
	"GoRestfulApi/model/domain"
	"context"
	"database/sql"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	return domain.Category{}
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	return domain.Category{}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, category int) domain.Category {
	return domain.Category{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	return []domain.Category{}
}
