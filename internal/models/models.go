package models

type User struct {
	Id               string `json:"user_id" db:"user_id"`
	Name             string `json:"user_name" db:"user_name"`
	PasswordHash     string `json:"password_hash" db:"password_hash"`
	AccessToken      string `json:"access_token" db:"access_token"`
	RefreshTokenHash string `json:"refresh_token_hash" db:"refresh_token_hash"`
}

type Route struct {
	Id           string  `json:"route_id" db:"route_id"`
	Speed        float32 `json:"route_speed" db:"route_speed"`
	AverageSpeed float32 `json:"route_average_speed" db:"route_average_speed"`
	Way          float32 `json:"route_way" db:"route_way"`
}

type RoutePoints struct {
	Id        string  `json:"cord_id" db:"cord_id"`
	RouteId   string  `json:"route_id" db:"route_id"`
	Width     float64 `json:"width" db:"width"`
	Longitude float64 `json:"longitude" db:"longitude"`
}
