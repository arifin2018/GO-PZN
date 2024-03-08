package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestServerFile(t *testing.T) {
	router := httprouter.New()
	router.ServeFiles("/src/*filepath", http.Dir("./resources"))

	request := httptest.NewRequest("GET", "http://localhost/src/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "hallo arifin", string(responseBody))
}

//go:embed resources
var resources embed.FS

func TestServerFile2(t *testing.T) {

	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "hallo arifin", string(body))
}
