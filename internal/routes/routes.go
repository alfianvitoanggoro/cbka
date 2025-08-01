package routes

import (
	"go-kafka/internal/factory"
	"go-kafka/pkg/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRouter(f *factory.Factory) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccess(w, http.StatusOK, "API Healthy", nil)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			RegisterAuthRoutes(r, f.AuthHandler)
		})

		r.Route("/users", func(r chi.Router) {
			RegisterUserRoutes(r, f.UserHandler)
		})
	})

	return r
}
