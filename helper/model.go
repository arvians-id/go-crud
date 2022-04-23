package helper

import (
	"go-crud/model/domain"
	"go-crud/model/web"
)

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
	}
}
