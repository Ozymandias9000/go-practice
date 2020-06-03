package domain

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Username        string `json:"username"`
}

func (d *Domain) Register(payload RegisterPayload) (*User, error) {
	userExists, err := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExists != nil {
		return nil, ErrUserAlreadyExists
	}

	userExists, err = d.DB.UserRepo.GetByUsername(payload.Username)
	if userExists != nil {
		return nil, ErrUserAlreadyExists
	}

	password, err := d.setPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: *password,
	}

	user, err := d.DB.UserRepo.CreateUser(&data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) setPassword(password string) (*string, error) {
	return &password, nil
}
