package routes

import (
	"GoResfulApiPribadi/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(categoryController controllers.CategoryController) {
	router := httprouter.New()
	router.GET("/", categoryController.Get)
	router.POST("/", categoryController.Insert)
	router.PUT("/:id", categoryController.Update)

	http.ListenAndServe(":8080", router)

}
