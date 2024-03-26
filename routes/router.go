package routes

import (
	"GoResfulApiPribadi/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router() {
	router := httprouter.New()
	router.GET("/", controllers.CategoryController)

	http.ListenAndServe(":8080", router)

}
