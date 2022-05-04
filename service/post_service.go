package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-crud/helper"
	"go-crud/model/domain"
	"go-crud/model/web"
	"go-crud/repository"
)

type PostService interface {
	FindAll(ctx context.Context) ([]web.PostResponse, error)
	FindById(ctx context.Context, postId int) (web.PostResponse, error)
	Create(ctx context.Context, request web.PostCreateRequest) (web.PostResponse, error)
	Update(ctx context.Context, request web.PostUpdateRequest) (web.PostResponse, error)
	Delete(ctx context.Context, postId int) error
}

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostServiceImpl(postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) *PostServiceImpl {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service PostServiceImpl) FindAll(ctx context.Context) ([]web.PostResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []web.PostResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	posts, err := service.PostRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.PostResponse{}, err
	}

	var postResponse []web.PostResponse
	for _, post := range posts {
		postResponse = append(postResponse, helper.ToPostResponse(post))
	}

	return postResponse, nil
}

func (service PostServiceImpl) FindById(ctx context.Context, postId int) (web.PostResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return web.PostResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		return web.PostResponse{}, err
	}

	return helper.ToPostResponse(rows), nil
}

func (service PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) (web.PostResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.PostResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.PostResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		Title:       request.Title,
		Description: request.Description,
	}

	post, err = service.PostRepository.Save(ctx, tx, post)
	if err != nil {
		return web.PostResponse{}, err
	}

	return helper.ToPostResponse(post), nil
}

func (service PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) (web.PostResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.PostResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.PostResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.PostResponse{}, err
	}

	rows.Title = request.Title
	rows.Description = request.Description

	rows, err = service.PostRepository.Update(ctx, tx, rows)
	if err != nil {
		return web.PostResponse{}, err
	}

	return helper.ToPostResponse(rows), nil
}

func (service PostServiceImpl) Delete(ctx context.Context, postId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		return err
	}

	err = service.PostRepository.Delete(ctx, tx, rows)
	if err != nil {
		return err
	}

	return nil
}
