package main

import (
	"GoResfulApiPribadi/controllers"
	"GoResfulApiPribadi/repositories"
	"GoResfulApiPribadi/routes"
	"GoResfulApiPribadi/services"
)

func main() {
	// DB := app.Database()

	categoryRepository := repositories.CategoryRepositoryConstruct()
	categoryServices := services.CategoryServiceConstruct(categoryRepository)
	categoryController := controllers.CategoryConstruct(categoryServices)
	// categoryController2 := con /

	routes.Router(categoryController)
}
