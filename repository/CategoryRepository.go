package repository

import "gopzn/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}