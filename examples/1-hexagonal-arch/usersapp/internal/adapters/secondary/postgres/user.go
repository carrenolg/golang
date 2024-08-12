package postgres

import (
	"context"
	"database/sql"
	"usersapp/internal/core/domains/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}

func (r *UserRepository) Add(ctx context.Context, u user.User) error {
	return nil
}
