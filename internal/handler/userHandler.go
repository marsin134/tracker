package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"tracker/internal/service"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		WriteErrorResponse(w, "Incorrect data", http.StatusBadRequest)
		return
	}

	response, err := h.service.User.Register(r.Context(), &req)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req service.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validate.Struct(req); err != nil {
		WriteErrorResponse(w, "Incorrect data", http.StatusBadRequest)
		return
	}

	response, err := h.service.User.Login(r.Context(), &req)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	userId := pathParts[len(pathParts)-1]

	user, err := h.service.User.GetUserById(r.Context(), userId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	userName := pathParts[len(pathParts)-1]

	user, err := h.service.User.GetUserByUsername(r.Context(), userName)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h Handler) UpdateRefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		WriteErrorResponse(w, "Session ID not found in request context", http.StatusBadRequest)
		return
	}

	response, err := h.service.User.UpdateRefreshToken(r.Context(), userId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		WriteErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		WriteErrorResponse(w, "Session ID not found in request context", http.StatusBadRequest)
		return
	}

	err := h.service.User.DeleteUser(r.Context(), userId)
	if err != nil {
		WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "Successful delete"})
}
