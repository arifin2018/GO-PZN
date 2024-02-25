package gopzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Data struct {
	Name string
	Grade string
}

func SimpleActionIf(writer http.ResponseWriter, request *http.Request) {
	t, e := template.ParseFiles("./templates/if.html")
	if e != nil {
		panic(e)
	}
	t.ExecuteTemplate(writer, "if.html", &Data{
		Name: "Arifin",
	})
}

func TestSimpleActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	SimpleActionIf(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func SimpleConditionOperator(writer http.ResponseWriter, request *http.Request) {
	t, e := template.ParseFiles("./templates/IfConditionOperator.html")
	if e != nil {
		panic(e)
	}
	t.ExecuteTemplate(writer, "IfConditionOperator.html", &Data{
		Name: "Arifin",
		Grade: "A",
	})
}

func TestSimpleConditionOperator(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	SimpleConditionOperator(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t, e := template.ParseFiles("./templates/loopRange.html")
	if e != nil {
		panic(e)
	}
	t.ExecuteTemplate(writer, "loopRange.html", map[string]any{
		"title":"arifin",
		"Hobbies":[]string{
			"catur","bola","coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateActionRange(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	t, e := template.ParseFiles("./templates/with.html")
	if e != nil {
		panic(e)
	}
	t.ExecuteTemplate(writer, "with.html", map[string]any{
		"Name":"arifin",
		"Address":map[string]string{
			"Street":"Gandaria",
			"City":"Jakarta",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateWith(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}