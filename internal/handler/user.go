// internal/handler/reconcile.go
package handler

import (
	"encoding/json"
	"net/http"

	"go-kafka/internal/dto"
	"go-kafka/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (h *UserHandler) UserReconcile(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestUserReconcile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.TriggeredBy == "" {
		http.Error(w, "triggered_by is required", http.StatusBadRequest)
		return
	}

	if err := h.usecase.UserReconcile(req.TriggeredBy); err != nil {
		http.Error(w, "failed to trigger reconcile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reconcile triggered successfully"))
}
