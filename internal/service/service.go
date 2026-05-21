package service

import (
	"context"
	"tracker/internal/config"
	"tracker/internal/models"
	"tracker/internal/repository"
)

type Service struct {
	User UserService
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	user := NewUserService(repo, cfg)
	return &Service{User: user}
}

type UserService interface {
	Register(ctx context.Context, req *UserRequest) (*UserResponse, error)
	Login(ctx context.Context, req *UserRequest) (*UserResponse, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateRefreshToken(ctx context.Context, userId string) (*UserResponse, error)
	DeleteUser(ctx context.Context, userId string) error
}
