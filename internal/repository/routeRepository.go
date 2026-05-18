package repository

import (
	"context"
	"tracker/internal/database"
	"tracker/internal/models"
)

type routeRepository struct {
	db *database.DB
}

func NewRouteRepository(db *database.DB) *routeRepository {
	return &routeRepository{db: db}
}

func (r routeRepository) CreateRoute(ctx context.Context, route *models.Route) (*string, error) {
	query := `INSERT INTO routes (route_id, user_id, route_speed, route_average_speed, route_way)
			VALUES (:route_id, :user_id, :route_speed, :route_average_speed, :route_way)`

	_, err := r.db.NamedExecContext(ctx, query, route)
	if err != nil {
		return nil, err
	}
	return &route.Id, nil
}

func (r routeRepository) GetRoute(ctx context.Context, id string) (*models.Route, error) {
	query := `SELECT * FROM routes WHERE route_id=$1`

	var route models.Route
	err := r.db.GetContext(ctx, &route, query, id)
	if err != nil {
		return nil, err
	}
	return &route, nil
}

func (r routeRepository) GetUserRoutes(ctx context.Context, userId string) (*[]models.Route, error) {
	query := `SELECT * FROM routes WHERE user_id=$1`

	var routes []models.Route
	err := r.db.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}
	return &routes, nil
}

func (r routeRepository) UpdateRoute(ctx context.Context, route *models.Route) (*string, error) {
	query := `UPDATE routes
		SET route_speed = :route_speed, route_average_speed = :route_average_speed, route_way = :route_way
    	WHERE route_id = :route_id`

	_, err := r.db.NamedExecContext(ctx, query, route)
	if err != nil {
		return nil, err
	}
	return &route.Id, nil
}

func (r routeRepository) DeleteRoute(ctx context.Context, id string) error {
	query := `DELETE FROM routes WHERE route_id=$1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
