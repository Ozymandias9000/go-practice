package handlers

import (
	"go-todo/domain"
	"net/http"
)

type authResponse struct {
	User  *domain.User `json:"user"`
	Token *domain.JWT  `json:"token"`
}

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := r.Context().Value("payload").(domain.RegisterPayload)

		user, err := s.domain.Register(payload)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		token, err := user.GenerateJWT()
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &authResponse{
			User:  user,
			Token: token,
		}, http.StatusCreated)
	}
}

func (s *Server) getCurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("currentUser")

		jsonResponse(w, user, http.StatusOK)
	}
}
