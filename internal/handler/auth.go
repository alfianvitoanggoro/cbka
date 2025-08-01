package handler

import (
	"encoding/json"
	"go-kafka/internal/dto"
	"go-kafka/internal/usecase"
	"go-kafka/pkg/response"
	"net/http"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestRegistration
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	createdUser, err := h.usecase.Create(&req)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response.WriteSuccess(w, http.StatusCreated, "User created successfully", createdUser)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	res, err := h.usecase.Login(&req)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to login", err.Error())
		return
	}

	response.WriteSuccess(w, http.StatusCreated, "User created successfully", res)
}
