package handlers

import (
	"context"
	"encoding/json"
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

func validatePayload(next http.HandlerFunc, payload interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		ctx := context.WithValue(r.Context(), "payload", payload)

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
