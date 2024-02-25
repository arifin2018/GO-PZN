package gopzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Header": "Arifin header",
		"H1":     "Arifin H1",
		"p": map[string]string{
			"name": "arifin",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateDataMap(response, *request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

type paraghraf struct {
	Name interface{}
}

type Datas struct {
	Header string
	H1     interface{}
	P      paraghraf
}

func TemplateDataStruct(writer http.ResponseWriter, request http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	Data := &Datas{
		Header: "arifin header",
		H1:     "arifin H1",
		P: paraghraf{
			Name: "arifin",
		},
	}
	fmt.Println(Data)
	t.ExecuteTemplate(writer, "name.gohtml", Data)
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateDataStruct(response, *request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}
