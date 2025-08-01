package routes

import (
	"go-kafka/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router, h *handler.AuthHandler) {
	r.Post("/registration", h.Registration)
	r.Post("/login", h.Login)
}
