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
		r.Route("/todos", func(r chi.Router) {
			r.Use(s.withUser)
			r.Get("/", s.getAllTodos())
			r.With(validatePayload(&domain.CreateTodoPayload{})).Post("/", s.createTodo())
			r.Route("/{id}", func(r chi.Router) {
				r.Use(s.withTodo)

				r.Get("/", s.getTodoByID())
				r.With(validatePayload(&domain.UpdateTodoPayload{})).Patch("/", s.updateTodo())
				r.Delete("/", s.deleteTodo())
			})
		})
	})
}
