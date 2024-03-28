package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Insert(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}
