// internal/handler/reconcile.go
package handler

import (
	"encoding/json"
	"net/http"

	"go-kafka/internal/dto"
	"go-kafka/internal/usecase"
	"go-kafka/pkg/response"
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
		response.WriteError(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if req.TriggeredBy == "" {
		response.WriteError(w, http.StatusBadRequest, "TriggeredBy is required", "TriggeredBy cannot be empty")
		return
	}

	if err := h.usecase.UserReconcile(req.TriggeredBy); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to trigger reconcile", err.Error())
		return
	}

	response.WriteSuccess(w, http.StatusOK, "Reconcile triggered successfully", nil)
}
