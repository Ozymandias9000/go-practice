package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type JWT struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func (u *User) GenerateJWT() (*JWT, error) {
	expiry := time.Now().Add(time.Hour)

	jwt := jwt.NewWithClaims(
		jwt.GetSigningMethod("HS256"),
		jwt.MapClaims{
			"id":  u.ID,
			"exp": expiry.Unix(),
		},
	)

	token, err := jwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JWT{
		AccessToken: token,
		ExpiresAt:   expiry,
	}, nil
}
