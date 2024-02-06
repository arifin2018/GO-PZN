package gopzn

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T)  {
	serve := http.Server{
		Addr: "localhost:2942",
	}
	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}