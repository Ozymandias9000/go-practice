package handlers

import (
	"encoding/json"
	"fmt"
	"go-todo/domain"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

/**
* Base response
 */
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type:", "application/json")

	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"data": data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

/**
* Error responses
 */
func UnauthorizedResponse(w http.ResponseWriter) {
	response := map[string]string{"error": domain.ErrUnauthorized.Error()}
	jsonResponse(w, response, http.StatusUnauthorized)
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
