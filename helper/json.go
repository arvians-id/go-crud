package helper

import (
	"encoding/json"
	"go-crud/model/web"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, request interface{}) {
	err := json.NewDecoder(r.Body).Decode(request)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := web.WebResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}

	err := json.NewEncoder(w).Encode(response)
	PanicIfError(err)
}
