package routes

import (
	"GoRestfulApi/controller"
	"GoRestfulApi/exception"
	"GoRestfulApi/helper"
	"GoRestfulApi/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(categoryController controller.CategoryController) {
	router := httprouter.New()

	Cors(router)
	router.NotFound = http.HandlerFunc(exception.CustomNotFound)
	router.MethodNotAllowed = http.HandlerFunc(exception.MethodNotAllowed)
	router.PanicHandler = exception.ErrorHandler

	router.GET("/api/categories", middleware.AuthorizeMiddleware(categoryController.FindAll))
	router.GET("/api/categories/:categoryId", middleware.AuthorizeMiddleware(categoryController.FindById))
	router.POST("/api/categories", middleware.AuthorizeMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId", middleware.AuthorizeMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId", middleware.AuthorizeMiddleware(categoryController.Delete))

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
