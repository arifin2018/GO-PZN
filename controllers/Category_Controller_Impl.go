package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
}

func CategoryConstruct() CategoryController {
	return &CategoryControllerImpl{}
}

func (CategoryControllerImpl *CategoryControllerImpl) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hai")
}

func (CategoryControllerImpl *CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (CategoryControllerImpl *CategoryControllerImpl) Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (CategoryControllerImpl *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (CategoryControllerImpl *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
