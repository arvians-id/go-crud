package middleware

import (
	"go-crud/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var apikey string = "TEKNIKBANG"
	if apikey == r.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		helper.WriteToResponseBody(w, http.StatusUnauthorized, nil)
	}
}
