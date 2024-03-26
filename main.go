package main

import (
	"GoResfulApiPribadi/controllers"
	"GoResfulApiPribadi/routes"
)

func main() {
	// DB := app.Database()

	// categoryRepository := repositories.CategoryRepository()
	categoryController := controllers.CategoryConstruct()
	// categoryController2 := con /

	routes.Router(categoryController)
}
