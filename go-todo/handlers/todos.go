package handlers

import (
	"go-todo/domain"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type todoResponse struct {
	Todo *domain.Todo `json:"todo"`
}

func (s *Server) createTodo() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := r.Context().Value("payload").(*domain.CreateTodoPayload)
		currentUser := r.Context().Value("currentUser").(*domain.User)
		currentUserID := currentUser.ID

		todo, err := s.domain.CreateTodo(payload, currentUserID)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &todoResponse{
			Todo: todo,
		}, http.StatusCreated)
	})
}

func (s *Server) getTodoByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		todoID, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		todo, err := s.domain.GetTodoByID(todoID)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &todoResponse{
			Todo: todo,
		}, http.StatusOK)
	})
}

func (s *Server) updateTodo() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := r.Context().Value("payload").(*domain.UpdateTodoPayload)
		todo := r.Context().Value("todo").(*domain.Todo)

		todo, err := s.domain.UpdateTodo(*payload, *todo)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &todoResponse{
			Todo: todo,
		}, http.StatusOK)
	})
}

func (s *Server) deleteTodo() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todo := r.Context().Value("todo").(*domain.Todo)

		todo, err := s.domain.DeleteTodo(*todo)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, &todoResponse{
			Todo: todo,
		}, http.StatusOK)
	})
}
