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