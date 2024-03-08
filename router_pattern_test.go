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

func TestRouterParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/image/*file", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("file")
		fmt.Fprintf(writer, "file: %v", image)
	})

	req, _ := http.NewRequest("GET", "/image/sampleid.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "file: /sampleid.png", string(body))
}
