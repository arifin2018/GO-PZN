package tests

import (
	"GoRestfulApi/controller"
	"GoRestfulApi/repositories"
	"GoRestfulApi/routes"
	"GoRestfulApi/services"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	db := NewDBTesting()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryServices := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryServices)
	return routes.Router(categoryController)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"name":"gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	bytes, _ := io.ReadAll(response.Body)
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, 200, responseBody["Code"])
	assert.Equal(t, "OK", responseBody["Status"])
}

func TestCreateCategoryFailed(t *testing.T) {

}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {
	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	bytes, _ := io.ReadAll(response.Body)
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, 200, int(responseBody["Code"].(float64)))
	assert.Equal(t, "gadget", responseBody["Data"].(map[string]interface{})["Name"])
}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestUnauthorize(t *testing.T) {

}
