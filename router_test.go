package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello world")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello world", string(body))
}

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
		fmt.Fprintf(w, "hello %v", p.ByName("name"))
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello/arifin", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "hello arifin", string(body))
}
