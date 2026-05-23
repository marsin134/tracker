package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tracker/internal/models"
)

type RoutePointRequest struct {
	RouteId   string  `json:"route_id" db:"route_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

func (h Handler) CreatePoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req RoutePointRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		WriteErrorResponse(w, "Incorrect data", http.StatusBadRequest)
		return
	}

	reqInModel := models.RoutePoints{
		RouteId:   req.RouteId,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	id, err := h.service.RoutePoint.CreatePoint(r.Context(), &reqInModel)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Route point created with id: %d", *id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: message})
}

func (h Handler) GetPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	pointId := pathParts[len(pathParts)-1]

	intId, err := strconv.ParseInt(pointId, 10, 64)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	point, err := h.service.RoutePoint.GetPoint(r.Context(), intId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(point)
}

func (h Handler) GetRoutePoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	routeId := pathParts[len(pathParts)-1]

	points, err := h.service.RoutePoint.GetRoutePoints(r.Context(), routeId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

func (h Handler) GetLastTwoPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	routeId := pathParts[len(pathParts)-1]

	points, err := h.service.RoutePoint.GetLastTwoPoints(r.Context(), routeId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

func (h Handler) DeletePoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	pathParts := strings.Split(r.URL.Path, "/")
	pointId := pathParts[len(pathParts)-1]
	intId, err := strconv.ParseInt(pointId, 10, 64)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.RoutePoint.DeletePoint(r.Context(), intId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "Successful delete"})
}
