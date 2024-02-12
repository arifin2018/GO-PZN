package gopzn

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	var cookie http.Cookie
	cookie.Name = "NurArifin"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, &cookie)
	fmt.Fprintln(writer, "Sukses membuat cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request)  {
	cookie, err := request.Cookie("NurArifin")
	if err != nil {
		panic(err)
	}else if cookie == nil {
		fmt.Fprintln(writer,"cookie kosong")
	}else{
		fmt.Fprintf(writer,"berhasil mendapatkan cookie anda : %s", cookie.Name)
	}
}

func TestCookie(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", SetCookie)
	mux.HandleFunc("/get", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}