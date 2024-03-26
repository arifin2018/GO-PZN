package repositories

import "context"

type CategoryRepository interface {
	Get(ctx context.Context)
	GetById()
	Insert()
	Update()
	Delete()
}
