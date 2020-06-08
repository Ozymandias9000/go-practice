package handlers

import (
	"go-todo/domain"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
