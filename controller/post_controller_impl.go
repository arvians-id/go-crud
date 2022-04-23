package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-crud/helper"
	"go-crud/model/web"
	"go-crud/service"
	"net/http"
	"strconv"
)

type PostControllerImpl struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (controller PostControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := controller.PostService.FindAll(r.Context())
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   post,
	}

	helper.WriteToResponseBody(w, response)
}

func (controller PostControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getParams := params.ByName("postId")
	id, err := strconv.Atoi(getParams)
	helper.PanicIfError(err)

	rows := controller.PostService.FindById(r.Context(), id)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   rows,
	}

	helper.WriteToResponseBody(w, response)
}

func (controller PostControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postCreateRequest := web.PostCreateRequest{}
	helper.ReadFromRequestBody(r, postCreateRequest)

	result := controller.PostService.Create(r.Context(), postCreateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}

func (controller PostControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postUpdateRequest := web.PostUpdateRequest{}
	helper.ReadFromRequestBody(r, postUpdateRequest)

	id, err := strconv.Atoi(params.ByName("postId"))
	helper.PanicIfError(err)
	postUpdateRequest.Id = id

	result := controller.PostService.Update(r.Context(), postUpdateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}

func (controller PostControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("postId"))
	helper.PanicIfError(err)
	controller.PostService.Delete(r.Context(), id)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, response)
}
