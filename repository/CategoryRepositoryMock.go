package repository

import (
	"fmt"
	"gopzn/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}else{
		category := arguments.Get(0).(entity.Category)
		fmt.Println(arguments.Get(0))
		return &category
	}
}