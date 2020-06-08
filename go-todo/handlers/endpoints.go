package handlers

import (
	"go-todo/domain"

	"github.com/go-chi/chi"
)

func (s Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.With(validatePayload(&domain.RegisterPayload{})).Post("/register", s.registerUser())
			r.With(s.withUser).Get("/currentUser", s.getCurrentUser())
		})
	})
}
