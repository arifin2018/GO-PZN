package gopzn

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var templates2 embed.FS

var myTemplates = template.Must(template.ParseFS(templates2, "templates/*.html"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "simple.html","hello HTML template")
}


func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateCaching(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}
