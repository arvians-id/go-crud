package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-crud/helper"
	"go-crud/model/domain"
)

type PostRepositoryImpl struct {
}

func NewPostRepositoryImpl() *PostRepositoryImpl {
	return &PostRepositoryImpl{}
}

func (repository PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	sqlQuery := "SELECT * FROM posts"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}

	return posts
}

func (repository PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	sqlQuery := "SELECT * FROM posts WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, postId)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var post domain.Post
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		helper.PanicIfError(err)
		return post, nil
	}

	return post, errors.New("post not found")
}

func (repository PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	sqlQuery := "INSERT INTO posts(title, description) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, sqlQuery, post.Title, post.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)

	return post
}

func (repository PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	sqlQuery := "UPDATE posts SET title = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, post.Title, post.Description, post.Id)
	helper.PanicIfError(err)

	return post
}

func (repository PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	sqlQuery := "DELETE FROM posts WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, post.Id)
	helper.PanicIfError(err)
}
