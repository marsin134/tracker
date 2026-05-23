package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h Handler) CreateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		WriteErrorResponse(w, "Session ID not found in request context", http.StatusBadRequest)
		return
	}

	id, err := h.service.Route.CreateRoute(r.Context(), userId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Route created with id: %s", *id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: message})
}

func (h Handler) GetRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	routeId := pathParts[len(pathParts)-1]

	fmt.Println(routeId)

	route, err := h.service.Route.GetRoute(r.Context(), routeId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(route)
}

func (h Handler) GetUserRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		WriteErrorResponse(w, "Session ID not found in request context", http.StatusBadRequest)
		return
	}

	routes, err := h.service.Route.GetUserRoutes(r.Context(), userId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

func (h Handler) UpdateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	routeId := pathParts[len(pathParts)-1]

	route, err := h.service.Route.UpdateRoute(r.Context(), routeId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(route)
}

func (h Handler) DeleteRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	routeId := pathParts[len(pathParts)-1]

	err := h.service.Route.DeleteRoute(r.Context(), routeId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "Successful delete"})
}
