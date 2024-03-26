package helper

import (
	Web_category "GoRestfulApi/model/web/Category"
	"encoding/json"
	"net/http"
)

func WebResponse(responseWriter http.ResponseWriter, code int, status string, err interface{}) {
	webResponse := Web_category.WebResponse{
		Code:   code,
		Status: status,
		Data:   err,
	}

	responseWriter.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(responseWriter)
	errors := encoder.Encode(webResponse)
	PanicIfError(errors)
}
