package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-crud/helper"
	"go-crud/model/web"
	"go-crud/service"
	"net/http"
	"strconv"
)

type PostController interface {
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type PostControllerImpl struct {
	PostService service.PostService
}

func NewPostControllerImpl(postService service.PostService) *PostControllerImpl {
	return &PostControllerImpl{PostService: postService}
}

func (controller PostControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	posts, err := controller.PostService.FindAll(r.Context())
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, posts)
}

func (controller PostControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getParams := params.ByName("postId")
	id, err := strconv.Atoi(getParams)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := controller.PostService.FindById(r.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusNotFound, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, post)
}

func (controller PostControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postCreateRequest := web.PostCreateRequest{}
	err := helper.ReadFromRequestBody(r, &postCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := controller.PostService.Create(r.Context(), postCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, post)
}

func (controller PostControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postUpdateRequest := web.PostUpdateRequest{}
	err := helper.ReadFromRequestBody(r, &postUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(params.ByName("postId"))
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}
	postUpdateRequest.Id = id

	post, err := controller.PostService.Update(r.Context(), postUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, post)
}

func (controller PostControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("postId"))
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = controller.PostService.Delete(r.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, nil)
}
