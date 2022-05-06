package helper

import (
	"encoding/json"
	"go-crud/model/web"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, request interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(request)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBody(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response := web.WebResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}

	err := json.NewEncoder(w).Encode(response)
	PanicIfError(err)
}
