package main

import (
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"go-crud/app"
	"go-crud/controller"
	"go-crud/helper"
	"go-crud/repository"
	"go-crud/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	postRepository := repository.NewPostRepositoryImpl()
	postService := service.NewPostServiceImpl(postRepository, db, validate)
	postController := controller.NewPostController(postService)

	router := httprouter.New()

	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:postId", postController.FindById)
	router.POST("/api/posts", postController.Create)
	router.PUT("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	fmt.Println("Running on localhost:8080...")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
