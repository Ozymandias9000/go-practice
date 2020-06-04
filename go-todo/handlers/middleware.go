package handlers

import (
	"context"
	"encoding/json"
	"go-todo/domain"
	"log"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

func validatePayload(next http.HandlerFunc, p interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload interface{}

		switch p.(type) {
		case domain.RegisterPayload:
			payload = &domain.RegisterPayload{}
		default:
			log.Println("No match found in validatePayload type switch")
		}

		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			badRequestResponse(w, err)
			return
		}

		v := validator.New()

		errs := v.Struct(payload)

		if errs != nil {
			validationErrorResponse(w, errs)
			return
		}

		defer r.Body.Close()

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
