package handlers

import (
	"User-service/db"
	"User-service/models"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Storage *db.User
}

func NewHandler(s *db.User) *Handler {
	return &Handler{Storage: s}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.UserRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.CreateNewUSer(body)
	if err != nil {
		http.Error(w, "failed creating new author", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
