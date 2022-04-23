package service

import (
	"context"
	"go-crud/model/web"
)

type PostService interface {
	FindAll(ctx context.Context) []web.PostResponse
	FindById(ctx context.Context, postId int) web.PostResponse
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse
	Delete(ctx context.Context, postId int)
}
