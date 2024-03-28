package controllers

import (
	"GoResfulApiPribadi/helpers"
	"GoResfulApiPribadi/model"
	"GoResfulApiPribadi/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func CategoryConstruct(service services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (CategoryControllerImpl *CategoryControllerImpl) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	CategoryControllerImpl.CategoryService.Get(r.Context())
}

func (CategoryControllerImpl *CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (CategoryControllerImpl *CategoryControllerImpl) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	helpers.PanicIfError(err)
	modelCategory := CategoryControllerImpl.CategoryService.Insert(r.Context(), &category)
	result := map[string]any{
		"Data": modelCategory,
	}
	helpers.JsonResponse(w, result)
}

func (CategoryControllerImpl *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	helpers.PanicIfError(err)
	number, err := strconv.Atoi(params.ByName("id"))
	helpers.PanicIfError(err)
	category.Id = number
	fmt.Println(category)
	modelCategory := CategoryControllerImpl.CategoryService.Update(r.Context(), &category)
	result := map[string]any{
		"Data": modelCategory,
	}
	helpers.JsonResponse(w, result)
}

func (CategoryControllerImpl *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
