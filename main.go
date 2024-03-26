package main

import (
	"GoResfulApiPribadi/app"
	"GoResfulApiPribadi/controllers"
	"GoResfulApiPribadi/repositories"
	"GoResfulApiPribadi/routes"
	"GoResfulApiPribadi/services"
)

func main() {
	DB := app.Database()

	categoryRepository := repositories.CategoryRepositoryConstruct()
	categoryServices := services.CategoryServiceConstruct(categoryRepository, DB)
	categoryController := controllers.CategoryConstruct(categoryServices)
	// categoryController2 := con /

	routes.Router(categoryController)
}
