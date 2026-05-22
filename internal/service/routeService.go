package service

import (
	"context"
	"github.com/google/uuid"
	"math"
	"tracker/internal/models"
	"tracker/internal/repository"
)

type routeService struct {
	repo *repository.Repository
}

func NewRouteService(repo *repository.Repository) *routeService {
	return &routeService{repo: repo}
}

func (svc routeService) CreateRoute(ctx context.Context, userId string) (*string, error) {
	route := models.Route{
		Id:           uuid.New().String(),
		UserId:       userId,
		Speed:        0,
		AverageSpeed: 0,
		Way:          0,
	}

	id, err := svc.repo.Route.CreateRoute(ctx, &route)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (svc routeService) GetRoute(ctx context.Context, id string) (*models.Route, error) {
	route, err := svc.repo.Route.GetRoute(ctx, id)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (svc routeService) GetUserRoutes(ctx context.Context, userId string) (*[]models.Route, error) {
	routes, err := svc.repo.Route.GetUserRoutes(ctx, userId)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (svc routeService) UpdateRoute(ctx context.Context, routeId string) (*models.Route, error) {
	route, err := svc.repo.Route.GetRoute(ctx, routeId)
	if err != nil {
		return nil, err
	}

	lastPoints, err := svc.repo.RoutePoint.GetLastTwoPoints(ctx, routeId)
	if err != nil {
		return nil, err
	}

	quantity, err := svc.repo.RoutePoint.GetRoutePoints(ctx, routeId)
	if err != nil {
		return nil, err
	}

	countSpeeds := len(*quantity) - 1

	pathTraveled := svc.calculatingMovement(*lastPoints)
	speed := svc.calculatingSpeed(pathTraveled, *lastPoints)
	averageSpeed := svc.calculatingAverageSpeed(route.AverageSpeed, speed, countSpeeds)

	route.Speed = speed
	route.AverageSpeed = averageSpeed
	route.Way += pathTraveled

	_, err = svc.repo.Route.UpdateRoute(ctx, route)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (svc routeService) DeleteRoute(ctx context.Context, id string) error {
	err := svc.repo.Route.DeleteRoute(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

const (
	radiusEarthEquator = 6378137.0
	squareEccentricity = 0.00669437999014
)

func (svc routeService) calculatingMovement(points []models.RoutePoints) float32 {
	// let's assign the corresponding values to the variables
	lat1 := points[0].Latitude
	lon1 := points[0].Longitude

	lat2 := points[1].Latitude
	lon2 := points[1].Longitude

	// Convert degrees to radians
	f1 := lat1 * math.Pi / 180
	f2 := lat2 * math.Pi / 180
	df := (lat2 - lat1) * math.Pi / 180
	dl := (lon2 - lon1) * math.Pi / 180

	// Average latitude
	fm := (f1 + f2) / 2
	sinFm := math.Sin(fm)

	// Meridian curvature radius (North-South)
	Rm := radiusEarthEquator * (1 - squareEccentricity) / math.Pow(1-squareEccentricity*sinFm*sinFm, 1.5)

	// Parallel curvature radius (East-West)
	Rn := radiusEarthEquator / math.Sqrt(1-squareEccentricity*sinFm*sinFm)

	// Offsets in meters
	dy := df * Rm
	dx := dl * Rn * math.Cos(fm)

	return float32(math.Sqrt(dx*dx + dy*dy))
}

func (svc routeService) calculatingSpeed(pathTraveled float32, points []models.RoutePoints) float32 {
	oldTime := points[1].CreatedAt
	newTime := points[0].CreatedAt

	timeElapsed := float32(newTime.Sub(oldTime).Seconds())
	return pathTraveled / timeElapsed
}

func (svc routeService) calculatingAverageSpeed(oldAverageSpeed float32, speed float32, quantity int) float32 {
	oldQuantity := float32(quantity - 1)
	return (oldAverageSpeed*oldQuantity + speed) / float32(quantity)
}
