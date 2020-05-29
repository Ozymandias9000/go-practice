package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	db     Database
	logger log.Logger
}

// NewService func
func NewService(db Database, logger log.Logger) Service {
	return &service{
		db,
		logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.db.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("error", err)
		return "", err
	}

	logger.Log("created user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.db.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("error", err)
		return "", err
	}

	logger.Log("got user", id)

	return email, nil
}
