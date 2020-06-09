package handlers

import (
	"context"
	"encoding/json"
	"go-todo/domain"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

func validate(w http.ResponseWriter, payload interface{}) error {
	v := validator.New()

	errs := v.Struct(payload)

	if errs != nil {
		validationErrorResponse(w, errs)
		return errs
	}
	return nil
}

func validatePayload(payload interface{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := json.NewDecoder(r.Body).Decode(&payload)

			if err != nil {
				badRequestResponse(w, err)
				return
			}

			defer r.Body.Close()

			err = validate(w, payload)
			if err != nil {
				return
			}

			p := payload

			ctx := context.WithValue(r.Context(), "payload", p)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (s *Server) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := s.domain.ParseToken(w, r)
		if err != nil {
			UnauthorizedResponse(w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			UnauthorizedResponse(w)
			return
		}

		id := int64(claims["id"].(float64))

		user, err := s.domain.DB.UserRepo.GetByID(id)

		ctx := context.WithValue(r.Context(), "currentUser", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) withTodo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo := new(domain.Todo)
		idString := chi.URLParam(r, "id")

		todoID, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		todo, err = s.domain.DB.TodoRepo.GetByID(todoID)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), "todo", todo)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
