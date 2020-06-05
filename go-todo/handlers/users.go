package handlers

import (
	"go-todo/domain"
	"net/http"
)

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.Context().Value("payload").(domain.RegisterPayload)

		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, user, http.StatusCreated)
	}
}
