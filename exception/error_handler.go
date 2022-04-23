package exception

import (
	"go-crud/helper"
	"go-crud/model/web"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, response)
}
