package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthorizeMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println("r.Response.Request.Body")
		log.Println(r.)
		n(w,r,p)
	}
}