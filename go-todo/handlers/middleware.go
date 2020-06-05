package handlers

import (
	"context"
	"encoding/json"
	"go-todo/domain"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

func validatePayload(next http.HandlerFunc, p interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		switch p.(type) {
		case domain.RegisterPayload:
			var payload domain.RegisterPayload

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

			ctx = context.WithValue(r.Context(), "payload", payload)
		default:
			log.Println("No match found in validatePayload type switch")
			badRequestResponse(w, domain.ErrWrongType)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (s *Server) withUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}

}
