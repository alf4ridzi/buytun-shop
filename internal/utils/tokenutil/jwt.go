package tokenutil

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtUserSession struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	jwt.RegisteredClaims
}

func CreateUserAccessToken(userID string, name string) (string, error) {
	expired := time.Now().Add(time.Duration(config.GetConfig().JwtAccessTokenExpired) * time.Hour)

	claims := &jwtUserSession{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetConfig().JwtAccessSecret))
}

func CreateUserRefreshToken(userID string) (string, error) {
	expired := time.Now().Add(time.Duration(config.GetConfig().JwtRefreshTokenExpired) * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expired),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetConfig().JwtRefreshSecret))
}

func ParseRefreshToken(tokenJwt string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}

	token, err := jwt.ParseWithClaims(tokenJwt, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error signing method")
		}

		return []byte(config.GetConfig().JwtRefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, domain.ErrJwtInvalidToken
	}

	return claims, nil
}
