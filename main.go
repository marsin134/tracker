package main

import (
	"context"
	"fmt"
	"tracker/internal/config"
	"tracker/internal/database"
	"tracker/internal/repository"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	logger := config.SetupLogger()
	// Uploading .env
	if err := config.LoadEnvFile(".env"); err != nil {
		logger.Errorf("Failed to load .env: %v", err)
	}

	cfg := config.LoadConfig()
	db, err := database.ConnectDB(&cfg, logger)
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo, err := repository.NewRepository(db)
	if err != nil {
		logger.Errorf("Failed to create repo: %v", err)
	}

	ctx := context.Background()

	err = repo.User.DeleteUser(ctx, "4403e218-740e-46cb-afe5-c5dbff5a4ddc")

	fmt.Println("GOOD")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
