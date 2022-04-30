package app

import (
	"github.com/julienschmidt/httprouter"
	"go-crud/controller"
	"go-crud/exception"
)

func NewRouter(postController controller.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:postId", postController.FindById)
	router.POST("/api/posts", postController.Create)
	router.PUT("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
