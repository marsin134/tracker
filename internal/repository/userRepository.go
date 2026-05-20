package repository

import (
	"context"
	"fmt"
	"time"
	"tracker/internal/database"
	"tracker/internal/models"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UpdateUserRefreshRequest struct {
	UserID       string
	RefreshToken string
	ExpiryTime   time.Time
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) (*string, error) {
	query := `
			INSERT INTO users (user_id, user_name, password_hash, refresh_token_hash, refresh_token_expiry_time)
			VALUES (:user_id, :user_name, :password_hash, :refresh_token_hash, :refresh_token_expiry_time)
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
	query := `SELECT * FROM users WHERE user_name=$1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) UpdateRefreshToken(ctx context.Context, req UpdateUserRefreshRequest) error {
	query := `
		UPDATE users 
		SET refresh_token_hash = $1, refresh_token_expiry_time = $2
		WHERE user_id = $3
	`

	_, err := r.db.ExecContext(ctx, query, req.RefreshToken, req.ExpiryTime, req.UserID)
	if err != nil {
		return fmt.Errorf("error updating refresh token: %w", err)
	}

	return nil
}

func (r UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE user_id=$1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
