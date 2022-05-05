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

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
		Image: user.Image,
	}
}
