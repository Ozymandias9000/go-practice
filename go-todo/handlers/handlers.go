package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-todo/domain"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type Server struct {
	domain *domain.Domain
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

func NewServer(d *domain.Domain) *Server {
	return &Server{domain: d}
}

func SetupRouter(d *domain.Domain) *chi.Mux {
	s := NewServer(d)

	r := chi.NewRouter()
	setupMiddlewares(r)

	s.setupEndpoints(r)

	return r
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type:", "application/json")

	w.WriteHeader(statusCode)

	if data == nil {
		data = map[string]string{}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func badRequestResponse(w http.ResponseWriter, err error) {
	response := map[string]string{"error": err.Error()}
	jsonResponse(w, response, http.StatusBadRequest)
}

func validationErrorResponse(w http.ResponseWriter, err error) {
	errResponse := make([]string, 0)

	for _, e := range err.(validator.ValidationErrors) {
		errResponse = append(errResponse, fmt.Sprint(e))
	}

	response := map[string][]string{"errors": errResponse}
	jsonResponse(w, response, http.StatusUnprocessableEntity)

}

func validatePayload(next http.HandlerFunc, payload *domain.RegisterPayload) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&payload)
		log.Printf("%v", payload)
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
