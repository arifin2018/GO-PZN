package gopzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.html", "./templates/footer.html", "./templates/layout.html"))
	t.ExecuteTemplate(writer, "layout.html", map[string]any{
		"Name": "arifin",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateLayout(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "BUDI"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "arifin",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateFunction(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len "BUDI"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "arifin",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateFunctionGlobal(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobalFunction(writer http.ResponseWriter, request *http.Request) {
	t,_ := template.New("FUNCTION").Funcs(map[string]interface{}{
		"upper":func (value string) string {
			return strings.ToUpper(value)
		},
	}).Parse(`{{ upper .Name }}`)
	


	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "arifin",
	})
}

func TestTemplateCreateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateCreateGlobalFunction(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobalFunctionPipeline(writer http.ResponseWriter, request *http.Request) {
	t,_ := template.New("FUNCTION").Funcs(map[string]interface{}{
		"SayHello":func (value string) string {
			return "Hello "+value
		},
		"upper":func (value string) string {
			return strings.ToUpper(value)
		},
	}).Parse(`{{ SayHello .Name | upper }}`)
	


	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "arifin",
	})
}

func TestTemplateCreateGlobalFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateCreateGlobalFunctionPipeline(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}
