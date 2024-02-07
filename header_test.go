package gopzn

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request)  {
	userAgent := request.Header.Get("Content-Type")
	fmt.Println(userAgent)
}

func TestRequestHeader(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080/hello",nil)
	request.Header.Add("Content-Type","application/json")
	writer := httptest.NewRecorder() 
	RequestHeader(writer, request)
}