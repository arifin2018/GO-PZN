package repositories

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Get(ctx context.Context, sqltx *sql.Tx)
	GetById()
	Insert()
	Update()
	Delete()
}
