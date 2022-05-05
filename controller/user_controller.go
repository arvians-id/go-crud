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
	var userRequest web.UserCreateRequest
	helper.ReadFromRequestBody(r, &userRequest)

	//uploadedFile, header, err := r.FormFile("image")
	//if err != nil {
	//	helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//defer uploadedFile.Close()
	//filename := header.Filename
	//
	//dir, err := os.Getwd()
	//if err != nil {
	//	helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//fileLocation := filepath.Join(dir, "assets/image", filename)
	//targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//defer targetFile.Close()
	//
	//_, err = io.Copy(targetFile, uploadedFile)
	//if err != nil {
	//	helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//userRequest.Image = filename

	user, err := controller.UserService.Create(r.Context(), userRequest)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteToResponseBody(w, http.StatusOK, user)
}

func (controller UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var userRequest web.UserUpdateRequest
	helper.ReadFromRequestBody(r, &userRequest)

	getParams := params.ByName("userId")
	id, err := strconv.Atoi(getParams)
	if err != nil {
		helper.WriteToResponseBody(w, http.StatusInternalServerError, err.Error())
		return
	}

	userRequest.Id = id

	user, err := controller.UserService.Update(r.Context(), userRequest)
	if err != nil {
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
