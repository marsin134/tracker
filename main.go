package main

import (
	"fmt"
	"tracker/internal/config"
	"tracker/internal/database"
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

	//repo, err := repository.NewRepository(db)
	//if err != nil {
	//	logger.Errorf("Failed to create repo: %v", err)
	//}
	//
	//ctx := context.Background()
	//
	//serv := service.NewService(repo, &cfg)
	//
	//point := &models.RoutePoints{
	//	RouteId:   "4eefafdb-b6a0-425e-867a-091f87eb616b",
	//	Latitude:  53.211933,
	//	Longitude: 44.969086,
	//}
	//
	//_, err = serv.RoutePoint.CreatePoint(ctx, point)
	//route, err := serv.Route.UpdateRoute(ctx, "4eefafdb-b6a0-425e-867a-091f87eb616b")
	//if err != nil {
	//	logger.Errorf("Failed to create route: %v", err)
	//}
	//fmt.Println(route)
	fmt.Println("GOOD")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
