package repository

import (
	"context"
	"database/sql"
	"errors"
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
	query := "SELECT * FROM users"
	queryContext, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.User{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var users []domain.User
	for queryContext.Next() {
		var user domain.User
		err := queryContext.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Image, &user.Password)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	queryContext, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return domain.User{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var user domain.User
	if queryContext.Next() {
		err := queryContext.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Image, &user.Password)
		if err != nil {
			return domain.User{}, err
		}

		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "INSERT INTO users (name,age,email,image,password) VALUES(?,?,?,?,?)"
	execContext, err := tx.ExecContext(ctx, query, user.Name, user.Age, user.Email, user.Image, user.Password)
	if err != nil {
		return domain.User{}, err
	}

	id, err := execContext.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	user.Id = int(id)

	return user, nil
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "UPDATE users SET name = ?, age = ?, email = ?, image = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.Age, user.Email, user.Image, user.Id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Id)
	if err != nil {
		return err
	}

	return nil
}
