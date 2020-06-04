package domain

import (
	"errors"
	"fmt"
)

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrNoResult          = errors.New("No result")
)

func ErrValidation(fields []string) error {
	msg := "Validation Error for the following fields: "
	for _, f := range fields {
		msg += fmt.Sprintf("%s, ", f)
	}
	return errors.New(msg)
}
