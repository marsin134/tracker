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

	//user := models.User{
	//	Id:               "521ce7f0-323c-4b12-a0b1-095b2221d5d6",
	//	Name:             "",
	//	PasswordHash:     "",
	//	AccessToken:      "",
	//	RefreshTokenHash: "",
	//}
	//
	//route := models.Route{
	//	Id:           "521ce7f0-323c-4b12-a0b1-095b2221d5d6",
	//	UserId:       user.Id,
	//	Speed:        0,
	//	AverageSpeed: 0,
	//	Way:          0,
	//}

	//point := models.RoutePoints{
	//	RouteId:   "521ce7f0-323c-4b12-a0b1-095b2221d5d6",
	//	Latitude:  0,
	//	Longitude: 0,
	//	CreatedAt: time.Now(),
	//}
	//
	//ctx := context.Background()

	//idUser, err := repo.User.CreateUser(ctx, &user)
	//if err != nil {
	//	logger.Errorf("Failed to create user: %v", err)
	//}
	//fmt.Printf("User created: %v", idUser)
	//
	//idRoute, err := repo.Route.CreateRoute(ctx, &route)
	//if err != nil {
	//	logger.Errorf("Failed to create route: %v", err)
	//}
	//fmt.Printf("Route created: %v", idRoute)

	//for i := 0; i < 5; i++ {
	//	id, err := repo.RoutePoint.CreatePoint(ctx, &point)
	//	if err != nil {
	//		logger.Errorf("Failed to create point: %v", err)
	//	}
	//	fmt.Printf("Created point ID: %d\n", id)
	//}
	//
	//points, err := repo.RoutePoint.GetLastTwoPoints(ctx, point.RouteId)
	//if err != nil {
	//	logger.Errorf("Failed to get last two points: %v", err)
	//}
	//fmt.Printf("Last two points: %v\n", *points)

	fmt.Println("GOOD")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
