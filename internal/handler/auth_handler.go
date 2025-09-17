package handler

import (
	"encoding/json"
	"net/http"

	"wallet/internal/entitty"
	"wallet/internal/usecase"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := &entitty.User{
		Username: req.Username,
		Password: req.Password,
	}

	id, err := h.uc.Register(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"id":       id,
		"username": req.Username,
		"status":   "created",
	})
}
