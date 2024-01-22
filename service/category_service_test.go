package service

import (
	"gopzn/entity"
	"gopzn/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceGetNotFound(t *testing.T){
	categoryRepository.Mock.On("FindById","1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceGetSuccess(t *testing.T){
	category := entity.Category{
		Id: "1",
		Name: "Arifin",
	}
	categoryRepository.Mock.On("FindById","2").Return(category)
	result, err := categoryService.Get("2")
	assert.Nil(t,err)
	assert.NotNil(t,result)
}