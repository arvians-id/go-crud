//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"go-crud/app"
	"go-crud/controller"
	"go-crud/middleware"
	"go-crud/repository"
	"go-crud/service"
	"net/http"
)

var postSet = wire.NewSet(
	repository.NewPostRepositoryImpl,
	wire.Bind(new(repository.PostRepository), new(*repository.PostRepositoryImpl)),
	service.NewPostServiceImpl,
	wire.Bind(new(service.PostService), new(*service.PostServiceImpl)),
	controller.NewPostControllerImpl,
	wire.Bind(new(controller.PostController), new(*controller.PostControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		postSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
