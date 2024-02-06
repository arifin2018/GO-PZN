package gopzn

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T)  {
	var Handler http.HandlerFunc = func (handle http.ResponseWriter,request *http.Request)  {
		fmt.Fprint(handle, "hello world")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: Handler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

func TestServeMux(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hi")
	})

	mux.HandleFunc("/arifin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "arifin")
	})

	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "image")
	})

	mux.HandleFunc("/image/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "image thumbnail")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}