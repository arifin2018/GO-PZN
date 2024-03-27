package services

import (
	"GoResfulApiPribadi/model"
	"context"
)

type CategoryService interface {
	Get(ctx context.Context)
	GetById()
	Insert(ctx context.Context, category *model.Category) *model.Category
	Update()
	Delete()
}
