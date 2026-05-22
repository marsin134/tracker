package service

import (
	"context"
	"tracker/internal/config"
	"tracker/internal/models"
	"tracker/internal/repository"
)

type Service struct {
	User       UserService
	Route      RouteService
	RoutePoint RoutePointService
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	user := NewUserService(repo, cfg)
	route := NewRouteService(repo)
	routePoints := NewRoutePointsService(repo)
	return &Service{
		User:       user,
		Route:      route,
		RoutePoint: routePoints,
	}
}

type UserService interface {
	Register(ctx context.Context, req *UserRequest) (*UserResponse, error)
	Login(ctx context.Context, req *UserRequest) (*UserResponse, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateRefreshToken(ctx context.Context, userId string) (*UserResponse, error)
	DeleteUser(ctx context.Context, userId string) error
}

type RouteService interface {
	CreateRoute(ctx context.Context, userId string) (*string, error)
	GetRoute(ctx context.Context, id string) (*models.Route, error)
	GetUserRoutes(ctx context.Context, userId string) (*[]models.Route, error)
	UpdateRoute(ctx context.Context, routeId string) (*models.Route, error)
	DeleteRoute(ctx context.Context, id string) error
}

type RoutePointService interface {
	CreatePoint(ctx context.Context, point *models.RoutePoints) (*int64, error)
	GetPoint(ctx context.Context, id int64) (*models.RoutePoints, error)
	GetRoutePoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error)
	DeletePoint(ctx context.Context, id int64) error
	GetLastTwoPoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error)
}
