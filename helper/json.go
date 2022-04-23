package helper

import (
	"encoding/json"
	"go-crud/model/web"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, postUpdateRequest interface{}) {
	err := json.NewDecoder(r.Body).Decode(&postUpdateRequest)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	PanicIfError(err)
}
