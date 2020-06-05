package handlers

import (
	"context"
	"encoding/json"
	"go-todo/domain"
	"log"
	"net/http"

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
