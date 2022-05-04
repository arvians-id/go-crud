package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-crud/model/domain"
)

type PostRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Post, error)
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error)
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error)
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post) error
}

type PostRepositoryImpl struct {
}

func NewPostRepositoryImpl() *PostRepositoryImpl {
	return &PostRepositoryImpl{}
}

func (repository PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Post, error) {
	sqlQuery := "SELECT * FROM posts"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	if err != nil {
		return []domain.Post{}, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			return []domain.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	sqlQuery := "SELECT * FROM posts WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, postId)
	if err != nil {
		return domain.Post{}, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var post domain.Post
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Description)
		if err != nil {
			return domain.Post{}, err
		}
		return post, nil
	}

	return post, errors.New("post not found")
}

func (repository PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error) {
	sqlQuery := "INSERT INTO posts(title, description) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, sqlQuery, post.Title, post.Description)
	if err != nil {
		return domain.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Post{}, err
	}

	post.Id = int(id)

	return post, nil
}

func (repository PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) (domain.Post, error) {
	sqlQuery := "UPDATE posts SET title = ?, description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, post.Title, post.Description, post.Id)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (repository PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) error {
	sqlQuery := "DELETE FROM posts WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlQuery, post.Id)
	if err != nil {
		return err
	}

	return nil
}
