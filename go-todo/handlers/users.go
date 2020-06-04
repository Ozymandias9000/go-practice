package handlers

import (
	"go-todo/domain"
	"net/http"
)

func (s *Server) registerUser() http.HandlerFunc {
	payload := domain.RegisterPayload{}
	return validatePayload(func(w http.ResponseWriter, r *http.Request) {

		// user, err := s.domain.Register(payload)
	}, &payload)

}
