package repositories

import (
	"GoResfulApiPribadi/model"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Get(ctx context.Context, sqltx *sql.Tx)
	GetById()
	Insert(ctx context.Context, sqltx *sql.Tx, category *model.Category) *model.Category
	Update()
	Delete()
}
