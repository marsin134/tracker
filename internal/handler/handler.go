package handler

import (
	"gitlab.com/golang-library/go-validator"
	"tracker/internal/config"
	"tracker/internal/service"
)

type Handler struct {
	service  *service.Service
	cfg      *config.Config
	Validate *validator.Validate
}

func NewHandler(service *service.Service, cfg *config.Config, validate *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		cfg:      cfg,
		Validate: validator.New(),
	}
}
