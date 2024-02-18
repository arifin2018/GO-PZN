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

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>templateText</title>
	</head>
	<body>
		{{ . }}
	</body>
	</html>`
	t,e :=template.New("SIMPLE").Parse(templateText)
	if e != nil{
		panic(e)
	}
	t.ExecuteTemplate(writer, "SIMPLE", "HELLO HTML TEMPLATE")
}

func TestSimpleHtml(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	response := httptest.NewRecorder()

	SimpleHtml(response,request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) {
	t,e := template.ParseFiles("./templates/simple.gohtml")
	if e != nil{
		panic(e)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "HELLO HTML TEMPLATE")
}

func TestSimpleHtmlFile(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	response := httptest.NewRecorder()

	SimpleHtmlFile(response,request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS
func SimpleHtmlEmbed(writer http.ResponseWriter, request *http.Request) {
	t,e := template.ParseFS(templates,"templates/*.gohtml")
	if e != nil{
		panic(e)
	}
	t.ExecuteTemplate(writer, "simple.gohtml", "HELLO HTML TEMPLATE")
}

func TestSimpleHtmlEmbed(t *testing.T)  {
	request := httptest.NewRequest("GET","http://localhost:8080",nil)
	response := httptest.NewRecorder()

	SimpleHtmlEmbed(response,request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}