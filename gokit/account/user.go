package account

import "context"

// User struct
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Database inter
type Database interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}
