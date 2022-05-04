package repository

import (
	"context"
	"database/sql"
	"go-crud/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) error
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) error {
	//TODO implement me
	panic("implement me")
}
