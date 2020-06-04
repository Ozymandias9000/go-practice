package handlers

import "go-todo/domain"

func validateRegisterUser(payload *domain.RegisterPayload) error {
	if payload.Password != payload.ConfirmPassword {
		fields := []string{"password", "confirmPassword"}
		return domain.ErrValidation(fields)
	}

	return nil
}
