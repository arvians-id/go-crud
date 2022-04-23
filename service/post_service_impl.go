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

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostServiceImpl(postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	var postResponse []web.PostResponse
	for _, post := range posts {
		postResponse = append(postResponse, helper.ToPostResponse(post))
	}

	return postResponse
}

func (service PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, postId)
	helper.PanicIfError(err)

	return helper.ToPostResponse(rows)
}

func (service PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		Title:       request.Title,
		Description: request.Description,
	}

	post = service.PostRepository.Save(ctx, tx, post)

	return helper.ToPostResponse(post)
}

func (service PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	rows.Title = request.Title
	rows.Description = request.Description

	rows = service.PostRepository.Update(ctx, tx, rows)

	return helper.ToPostResponse(rows)
}

func (service PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rows, err := service.PostRepository.FindById(ctx, tx, postId)
	helper.PanicIfError(err)

	service.PostRepository.Delete(ctx, tx, rows)
}
