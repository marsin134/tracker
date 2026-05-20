package repository

import (
	"context"
	"tracker/internal/database"
	"tracker/internal/models"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) (*string, error) {
	query := `
			INSERT INTO users (user_id, user_name, password_hash, access_token, refresh_token_hash)
			VALUES (:user_id, :user_name, :password_hash, :access_token, :refresh_token_hash)
			`
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return nil, err
	}
	return &user.Id, nil
}

func (r UserRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT * FROM users WHERE user_id=$1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username=$1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE user_id=$1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
