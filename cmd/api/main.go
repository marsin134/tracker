package main

import (
	"fmt"
	"log"
	"net/http"
	"tracker/cmd/app"
	"tracker/internal/config"
)

func main() {
	logger := config.SetupLogger()
	// Uploading .env
	if err := config.LoadEnvFile(".env"); err != nil {
		log.Printf("Failed to load .env: %v", err)
	}

	cfg := config.LoadConfig()

	if cfg.Token.JWTSecret == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the .env file")
	}

	db, svc := app.App(&cfg, logger)
	defer db.Close()

	handlerChain := app.InitializationHandlers(svc, &cfg)

	// Starting the server
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	fmt.Printf("The server is running on %s\n", addr)

	if err := http.ListenAndServe(addr, handlerChain); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}
}
