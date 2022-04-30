package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-crud/helper"
	"go-crud/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	fmt.Println("You're Running on localhost:8080")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
