package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrNoResult          = errors.New("No result")
)