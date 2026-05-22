package service

import (
	"context"
	"time"
	"tracker/internal/models"
	"tracker/internal/repository"
)

type routePointsService struct {
	repo *repository.Repository
}

func NewRoutePointsService(repo *repository.Repository) *routePointsService {
	return &routePointsService{repo: repo}
}

func (svc routePointsService) CreatePoint(ctx context.Context, point *models.RoutePoints) (*int64, error) {
	point.CreatedAt = time.Now()
	id, err := svc.repo.RoutePoint.CreatePoint(ctx, point)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (svc routePointsService) GetPoint(ctx context.Context, id int64) (*models.RoutePoints, error) {
	point, err := svc.repo.RoutePoint.GetPointById(ctx, id)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (svc routePointsService) GetRoutePoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error) {
	points, err := svc.repo.RoutePoint.GetRoutePoints(ctx, routeId)
	if err != nil {
		return nil, err
	}
	return points, nil
}

func (svc routePointsService) DeletePoint(ctx context.Context, id int64) error {
	err := svc.repo.RoutePoint.DeletePoint(ctx, id)
	return err
}

func (svc routePointsService) GetLastTwoPoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error) {
	points, err := svc.repo.RoutePoint.GetLastTwoPoints(ctx, routeId)
	if err != nil {
		return nil, err
	}
	return points, nil
}
