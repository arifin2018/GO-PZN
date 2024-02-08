package gopzn

import (
	"fmt"
	"net/http"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer,"Name is empty")
	}else{
		fmt.Fprint(writer, "Hello ", name)
	}
}