package controller

import (
	"GoRestfulApi/app"
	"GoRestfulApi/helper"
	Web_category "GoRestfulApi/model/web/Category"
	"GoRestfulApi/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (CategoryControllerImpl *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer app.NewDB().Close()
	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := Web_category.CreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := CategoryControllerImpl.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := Web_category.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (CategoryControllerImpl *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer app.NewDB().Close()
	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := Web_category.UpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := CategoryControllerImpl.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := Web_category.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (CategoryControllerImpl *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer app.NewDB().Close()
	defer request.Body.Close()
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	CategoryControllerImpl.CategoryService.Delete(request.Context(), id)
	webResponse := Web_category.WebResponse{
		Code:   200,
		Status: "OK",
	}
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (CategoryControllerImpl *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer app.NewDB().Close()
	defer request.Body.Close()
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("panic occurred:", err)
	// 	}
	// }()
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := CategoryControllerImpl.CategoryService.FindById(request.Context(), id)
	webResponse := Web_category.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (CategoryControllerImpl *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	defer app.NewDB().Close()
	defer request.Body.Close()
	categoryResponses := CategoryControllerImpl.CategoryService.FindAll(request.Context())
	webResponse := Web_category.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
