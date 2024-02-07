package gopzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Println(writer,"Hello")
	}else{
		fmt.Println(writer,"Hello ", name)
	}
}

func TestQueryParams(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080/hello?name=Arifin",nil)
	response := httptest.NewRecorder()
	
	SayHello(response, request)

	response2 := response.Result()
	data, _ := io.ReadAll(response2.Body)
	
	fmt.Println(data)
}

func MultipleSayHello(writer http.ResponseWriter, request *http.Request)  {
	firstname := request.URL.Query().Get("firstname")
	lastname := request.URL.Query().Get("lastname")
	if firstname == "" || lastname == "" {
		fmt.Println(writer,"Hello")
	}else{
		fmt.Println(writer,"Hello ", firstname, lastname)
	}
}

func TestMultipleSayHello(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080/hello?firstname=Nur&lastname=Arifin",nil)
	response := httptest.NewRecorder()
	
	MultipleSayHello(response, request)

	response2 := response.Result()
	data, _ := io.ReadAll(response2.Body)
	
	fmt.Println(data)
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request)  {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names," "))
}

func TestMultipleParameterValues(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080/hello?name=Nur&name=Arifin",nil)
	response := httptest.NewRecorder()
	
	MultipleParameterValues(response, request)

	response2 := response.Result()
	data, _ := io.ReadAll(response2.Body)
	
	fmt.Println(string(data))
}