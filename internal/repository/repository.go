package repository

import (
	"context"
	"tracker/internal/database"
	"tracker/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (*string, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type RouteRepo interface {
	CreateRoute(ctx context.Context, route *models.Route) (*string, error)
	GetRoute(ctx context.Context, id string) (*models.Route, error)
	GetUserRoutes(ctx context.Context, userId string) (*[]models.Route, error)
	UpdateRoute(ctx context.Context, route *models.Route) (*string, error)
	DeleteRoute(ctx context.Context, id string) error
}

type RoutePointsRepo interface {
	CreatePoint(ctx context.Context, point *models.RoutePoints) (*int64, error)
	GetPointById(ctx context.Context, id int64) (*models.RoutePoints, error)
	GetRoutePoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error)
	DeletePoint(ctx context.Context, id int64) error
	GetLastTwoPoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error)
}

type Repository struct {
	User       UserRepo
	Route      RouteRepo
	RoutePoint RoutePointsRepo
}

func NewRepository(db *database.DB) (*Repository, error) {
	user := NewUserRepository(db)
	route := NewRouteRepository(db)
	routePoint := NewRoutePointsRepository(db)
	return &Repository{User: user,
		Route:      route,
		RoutePoint: routePoint}, nil
}
