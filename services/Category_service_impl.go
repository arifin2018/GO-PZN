package services

import (
	"GoRestfulApi/helper"
	"GoRestfulApi/model/domain"
	Web_category "GoRestfulApi/model/web/Category"
	"GoRestfulApi/repositories"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(CategoryRepository repositories.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: CategoryRepository,
		DB:                 DB,
		validate:           validate,
	}
}

func (categoryServiceImpl *CategoryServiceImpl) Create(ctx context.Context, WebCategory Web_category.CreateRequest) Web_category.CategoryResponse {
	err := categoryServiceImpl.validate.Struct(WebCategory)
	helper.PanicIfError(err)
	tx, err := categoryServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: WebCategory.Name,
	}

	category = categoryServiceImpl.CategoryRepository.Save(ctx, tx, category)

	return Web_category.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (categoryServiceImpl *CategoryServiceImpl) Update(ctx context.Context, WebCategory Web_category.UpdateRequest) Web_category.CategoryResponse {
	err := categoryServiceImpl.validate.Struct(WebCategory)
	helper.PanicIfError(err)
	tx, err := categoryServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryServiceImpl.CategoryRepository.FindById(ctx, tx, WebCategory.Id)
	helper.PanicIfError(err)
	category.Name = WebCategory.Name
	category = categoryServiceImpl.CategoryRepository.Update(ctx, tx, category)

	return Web_category.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (categoryServiceImpl *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := categoryServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryServiceImpl.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	categoryServiceImpl.CategoryRepository.Delete(ctx, tx, category)
}

func (categoryServiceImpl *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) Web_category.CategoryResponse {
	tx, err := categoryServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryServiceImpl.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return Web_category.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (categoryServiceImpl *CategoryServiceImpl) FindAll(ctx context.Context) []Web_category.CategoryResponse {
	tx, err := categoryServiceImpl.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var categoriesResponse []Web_category.CategoryResponse
	categories := categoryServiceImpl.CategoryRepository.FindAll(ctx, tx)

	for _, v := range categories {
		categoriesResponse = append(categoriesResponse, Web_category.CategoryResponse{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return categoriesResponse
}
