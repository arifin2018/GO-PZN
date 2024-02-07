package gopzn

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request http.Request)  {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firsname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")

	fmt.Fprint(writer,firsname, lastname)
}

func TestFormPost(t *testing.T)  {
	requestBody := strings.NewReader("firstname=nur&lastname=arifin")
	request := httptest.NewRequest("POST","http://localhost:8080/hello",requestBody)
	request.Header.Add("Content-type","application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, *request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}