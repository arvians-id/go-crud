package repository

import (
	"context"
	"database/sql"
	"go-crud/model/domain"
)

type PostRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, postId int) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, postId int)
}
