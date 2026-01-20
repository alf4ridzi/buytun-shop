package tokenutil

import (
	"buytun-backend/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtUserSession struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func CreateUserAuthToken(userID string, name string) (string, error) {
	expired := time.Now().Add(time.Duration(config.GetConfig().JwtAccessTokenExpired) * time.Hour)

	claims := &jwtUserSession{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
