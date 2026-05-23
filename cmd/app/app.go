package app

import (
	"github.com/sirupsen/logrus"
	"tracker/internal/config"
	"tracker/internal/database"
	"tracker/internal/repository"
	"tracker/internal/service"
)

func App(cfg *config.Config, logger *logrus.Logger) (*database.DB, *service.Service) {
	db, err := database.ConnectDB(cfg, logger)
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
	}

	repo, err := repository.NewRepository(db)
	if err != nil {
		logger.Errorf("Failed to create repo: %v", err)
	}

	svc := service.NewService(repo, cfg)

	return db, svc
}
