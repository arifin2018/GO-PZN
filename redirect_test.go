package gopzn

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request)  {
	http.Redirect(writer,request,"/redirect_to", http.StatusPermanentRedirect)
}

func TestRedirect(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect_from", RedirectFrom)
	mux.HandleFunc("/redirect_to", RedirectTo)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}