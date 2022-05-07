package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-crud/helper"
	"go-crud/model/domain"
	"go-crud/model/web"
	"go-crud/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	FindById(ctx context.Context, userId int) (web.UserResponse, error)
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, userId int) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (service UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.UserResponse{}, err
	}

	var userResponse []web.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, helper.ToUserResponse(user))
	}

	return userResponse, nil
}

func (service UserServiceImpl) FindById(ctx context.Context, userId int) (web.UserResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}

func (service UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	createUser := domain.User{
		Name:     request.Name,
		Age:      request.Age,
		Email:    request.Email,
		Image:    request.Image,
		Password: string(password),
	}
	user, err := service.UserRepository.Save(ctx, tx, createUser)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}

func (service UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	getUser, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.UserResponse{}, err
	}

	newImage := request.Image
	oldImage := getUser.Image
	if newImage == "" {
		newImage = getUser.Image
	}

	getUser.Name = request.Name
	getUser.Age = request.Age
	getUser.Email = request.Email
	getUser.Image = newImage

	user, err := service.UserRepository.Update(ctx, tx, getUser)
	if err != nil {
		return web.UserResponse{}, err
	}

	if newImage != "" {
		_ = helper.DeleteImage("assets/images", oldImage)
	}

	return helper.ToUserResponse(user), nil
}

func (service UserServiceImpl) Delete(ctx context.Context, userId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	getUser, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return err
	}

	_ = helper.DeleteImage("assets/images", getUser.Image)

	err = service.UserRepository.Delete(ctx, tx, getUser)
	if err != nil {
		return err
	}

	return nil
}
