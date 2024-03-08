package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New() //untuk define router new
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello world")
	})
	// router.GET("/hello/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// 	fmt.Fprintf(writer, "hello %v", params.ByName("name"))
	// })

	// server := http.Server{
	// 	Handler: router,           //handler untuk router new
	// 	Addr:    "localhost:8080", //define address url
	// }

	// server.ListenAndServe() //untuk running server
	log.Fatal(http.ListenAndServe(":8080", router))
}
