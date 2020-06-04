package handlers

import (
	"encoding/json"
	"go-todo/domain"
	"net/http"
)

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := domain.RegisterPayload{}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		if errs := validateRegisterUser(&payload); errs != nil {
			validationErrorResponse(w, err)
			return
		}

		// user, err := s.domain.Register(payload)
	}
}
