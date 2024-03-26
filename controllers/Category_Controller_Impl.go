package controllers

import (
	"GoResfulApiPribadi/services"
	"net/http"

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

}

func (CategoryControllerImpl *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (CategoryControllerImpl *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
