package main

import (
	"GoRestfulApi/controller"
	"GoRestfulApi/repositories"
	"GoRestfulApi/services"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoriesId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoriesId", categoryController.Update)
	router.DELETE("/api/categories/:categoriesId", categoryController.Delete)
}
