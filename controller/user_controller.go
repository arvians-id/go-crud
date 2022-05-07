package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-crud/helper"
	"go-crud/model/web"
	"go-crud/service"
	"net/http"
	"strconv"
)

type UserController interface {
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users, err := controller.UserService.FindAll(r.Context())
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, users)
}

func (controller UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getParams := params.ByName("userId")
	id, err := strconv.Atoi(getParams)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := controller.UserService.FindById(r.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, user)
}

func (controller UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// File Upload
	path := "assets/images"
	filename, err := helper.UploadImage(r, 1, path)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	// User Request
	age, err := strconv.Atoi(r.PostFormValue("age"))
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	userRequest := web.UserCreateRequest{
		Name:     r.PostFormValue("name"),
		Age:      age,
		Email:    r.PostFormValue("email"),
		Image:    filename,
		Password: r.PostFormValue("password"),
	}

	user, err := controller.UserService.Create(r.Context(), userRequest)
	if err != nil {
		_ = helper.DeleteImage(path, filename)
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteToResponseBody(w, http.StatusOK, user)
}

func (controller UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// File Upload
	_, _, err := r.FormFile("image")
	var filename string
	path := "assets/images"
	// Is File Exists
	if err == nil {
		filename, err = helper.UploadImage(r, 1, path)
		if err != nil {
			helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// User Request
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	getParams := params.ByName("userId")
	id, err := strconv.Atoi(getParams)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	userRequest := web.UserUpdateRequest{
		Id:    id,
		Name:  r.FormValue("name"),
		Age:   age,
		Email: r.FormValue("email"),
		Image: filename,
	}

	user, err := controller.UserService.Update(r.Context(), userRequest)
	if err != nil {
		_ = helper.DeleteImage(path, filename)
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, user)
}

func (controller UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	getParams := params.ByName("userId")
	id, err := strconv.Atoi(getParams)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = controller.UserService.Delete(r.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.WriteToResponseBody(w, http.StatusOK, nil)
}
