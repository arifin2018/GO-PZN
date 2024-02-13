package gopzn

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	file:= http.Dir("./resources")
	fileserver := http.FileServer(file)
	fmt.Println(fileserver)

	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static",fileserver))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//go:embed resources
var resources embed.FS
func TestFileServer2(t *testing.T) {
	file,_:= fs.Sub(resources,"resources")
	fileserver := http.FileServer(http.FS(file))

	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static",fileserver))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}