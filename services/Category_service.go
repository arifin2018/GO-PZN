package services

import (
	Web_category "GoRestfulApi/model/web/Category"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, WebCategory Web_category.CreateRequest) Web_category.CategoryResponse
	Update(ctx context.Context, WebCategory Web_category.UpdateRequest) Web_category.CategoryResponse
	Delete(ctx context.Context, cateforyId int)
	FindById(ctx context.Context, cateforyId int) Web_category.CategoryResponse
	FindAll(ctx context.Context) []Web_category.CategoryResponse
}
