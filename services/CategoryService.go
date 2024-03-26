package services

import "context"

type CategoryService interface {
	Get(ctx context.Context)
	GetById()
	Insert()
	Update()
	Delete()
}
