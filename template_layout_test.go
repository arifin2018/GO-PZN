package gopzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
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
