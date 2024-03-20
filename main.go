package main

import (
	"GoRestfulApi/app"
	"GoRestfulApi/controller"
	"GoRestfulApi/repositories"
	"GoRestfulApi/routes"
	"GoRestfulApi/services"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	routes.Router(categoryController)
}
