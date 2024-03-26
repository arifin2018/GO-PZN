package main

import (
	"GoRestfulApi/app"
	"GoRestfulApi/controller"
	"GoRestfulApi/helper"
	"GoRestfulApi/repositories"
	"GoRestfulApi/routes"
	"GoRestfulApi/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := routes.Router(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
