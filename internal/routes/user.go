// internal/routes/reconcile.go
package routes

import (
	"go-kafka/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, h *handler.UserHandler) {
	r.Post("/reconcile", h.UserReconcile)
}
