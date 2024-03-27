package helpers

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, result map[string]any) {
	w.Header().Add("Content-Type", "Application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(result)
	PanicIfError(err)
}
