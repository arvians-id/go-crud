package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-crud/model/web"
	"go-crud/repository"
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

func (u UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) FindById(ctx context.Context, userId int) (web.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Delete(ctx context.Context, userId int) error {
	//TODO implement me
	panic("implement me")
}
