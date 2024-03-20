package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func headerSet(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
	}
	w.WriteHeader(http.StatusNoContent)
}

func Cors(router *httprouter.Router) {
	router.GlobalOPTIONS = http.HandlerFunc(headerSet)
}
