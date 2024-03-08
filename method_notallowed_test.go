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

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "method not allowed")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseResult := response.Result()
	body, _ := io.ReadAll(responseResult.Body)

	assert.Equal(t, "method not allowed", string(body))
}
