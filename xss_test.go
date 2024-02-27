package gopzn

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.html",map[string]interface{}{
		"Title":"Template Auto Escape",
		"Body":"<p>Ini adalah body</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateAutoEscape(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.html",map[string]interface{}{
		"Title":"Template Auto Escape",
		"Body": template.HTML("<p>Ini adalah body</p>") ,
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
