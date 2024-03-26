package exception

import (
	"GoRestfulApi/helper"
	"net/http"
)

func ErrorHandler(responseWriter http.ResponseWriter, request *http.Request, err interface{}) {
	responseWriter.Header().Set("Content-type", "application/json")
	responseWriter.WriteHeader(http.StatusInternalServerError)
	helper.WebResponse(responseWriter, http.StatusInternalServerError, "StatusInternalServerError", err)
}

func CustomNotFound(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")
	responseWriter.WriteHeader(http.StatusNotFound)
	helper.WebResponse(responseWriter, http.StatusNotFound, "StatusNotFound", nil)
}

func Unauthorize(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")
	responseWriter.WriteHeader(http.StatusUnauthorized)
	helper.WebResponse(responseWriter, http.StatusUnauthorized, "Unauthorize", nil)
}

func MethodNotAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	helper.WebResponse(responseWriter, http.StatusNotFound, "StatusMethodNotAllowed", nil)
}
