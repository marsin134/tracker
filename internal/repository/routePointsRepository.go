package repository

import (
	"context"
	"tracker/internal/database"
	"tracker/internal/models"
)

type RoutePointsRepository struct {
	db *database.DB
}

func NewRoutePointsRepository(db *database.DB) *RoutePointsRepository {
	return &RoutePointsRepository{db: db}
}

func (r RoutePointsRepository) CreatePoint(ctx context.Context, point *models.RoutePoints) (*int64, error) {
	query := `INSERT INTO route_points (route_id, latitude, longitude, created_at)
	VALUES (:route_id, :latitude, :longitude, :created_at)`

	_, err := r.db.NamedExecContext(ctx, query, point)
	if err != nil {
		return nil, err
	}

	query = `SELECT * FROM route_points WHERE route_id=$1 ORDER BY id DESC LIMIT 1`

	var points []models.RoutePoints
	err = r.db.SelectContext(ctx, &points, query, point.RouteId)
	if err != nil {
		return nil, err
	}

	return &points[0].Id, nil
}

func (r RoutePointsRepository) GetPointById(ctx context.Context, id int64) (*models.RoutePoints, error) {
	query := `SELECT * FROM route_points WHERE id=$1`

	var point models.RoutePoints
	err := r.db.GetContext(ctx, &point, query, id)
	if err != nil {
		return nil, err
	}
	return &point, nil
}

func (r RoutePointsRepository) GetRoutePoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error) {
	query := `SELECT * FROM route_points WHERE route_id=$1`

	var points []models.RoutePoints
	err := r.db.SelectContext(ctx, &points, query, routeId)
	if err != nil {
		return nil, err
	}
	return &points, nil
}

func (r RoutePointsRepository) DeletePoint(ctx context.Context, id int64) error {
	query := `DELETE FROM route_points WHERE id=$1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RoutePointsRepository) GetLastTwoPoints(ctx context.Context, routeId string) (*[]models.RoutePoints, error) {
	query := `SELECT * FROM route_points WHERE route_id=$1 ORDER BY id DESC LIMIT 2`
	var points []models.RoutePoints

	err := r.db.SelectContext(ctx, &points, query, routeId)
	if err != nil {
		return nil, err
	}
	return &points, nil
}
