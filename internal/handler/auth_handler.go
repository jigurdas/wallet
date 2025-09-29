package handler

import (
	"encoding/json"
	"net/http"

	"wallet/internal/entity"
	"wallet/internal/usecase"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	uc *usecase.AuthUsecase
}

func NewUserHandler(uc *usecase.AuthUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username     string `json:"username"`
		PasswordHash string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.PasswordHash == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}

	// Hash password before sending to usecase/repo
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	user := &entity.User{
		Username:     req.Username,
		PasswordHash: string(hashed),
	}

	id, err := h.uc.Register(r.Context(), user)
	if err != nil {
		http.Error(w, "failed to register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id":       id,
		"username": req.Username,
		"status":   "created",
	})
}
