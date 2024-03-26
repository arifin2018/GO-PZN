package middleware

import (
	"GoRestfulApi/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthorizeMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Header.Get("X-API-KEY") != "RAHASIA" {
			exception.Unauthorize(w, r)
			return
		} else {
			n(w, r, p)
		}
	}
}
