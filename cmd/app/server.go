package app

import (
	"fmt"
	"net/http"
	"tracker/internal/config"
	"tracker/internal/handler"
	"tracker/internal/middleware"
	"tracker/internal/service"
)

func InitializationHandlers(svc *service.Service, cfg *config.Config) http.Handler {
	handlers := handler.NewHandler(svc, cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)

	mux.HandleFunc("/api/auth/register", handlers.Register)
	mux.HandleFunc("/api/auth/login", handlers.Login)

	mux.HandleFunc("/api/user/id/", handlers.GetUserById)
	mux.HandleFunc("/api/user/name/", handlers.GetUserByUsername)
	mux.HandleFunc("/api/user/updateRefreshToken", handlers.UpdateRefreshToken)
	mux.HandleFunc("/api/user/delete/", handlers.DeleteUser)

	mux.HandleFunc("/api/route/create", handlers.CreateRoute)
	mux.HandleFunc("/api/route/get/", handlers.GetRoute)
	mux.HandleFunc("/api/route/get", handlers.GetUserRoutes)
	mux.HandleFunc("/api/route/update/", handlers.UpdateRoute)
	mux.HandleFunc("/api/route/delete/", handlers.DeleteRoute)

	mux.HandleFunc("/api/point/create", handlers.CreatePoint)
	mux.HandleFunc("/api/point/get/point/", handlers.GetPoint)
	mux.HandleFunc("/api/point/get/routePoints/", handlers.GetRoutePoints)
	mux.HandleFunc("/api/point/get/routeLastPoint/", handlers.GetLastTwoPoints)
	mux.HandleFunc("/api/point/delete/", handlers.DeletePoint)

	handlerChain := middleware.Chain(
		mux,
		middleware.CORSMiddleware,
		middleware.AuthMiddleware(cfg),
		middleware.LoggingMiddleware)

	return handlerChain
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "/, HomeHandler")
}
