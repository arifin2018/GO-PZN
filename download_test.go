package gopzn

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")
	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "%x\n", "Bad Request")
		return
	}
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"") // ini hanya untuk download file tanpa render,jika ingin merender,silahkan hapus ini tanpa download
	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
