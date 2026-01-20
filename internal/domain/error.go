package domain

import "errors"

var (
	// auth error
	ErrInvalidAuth = errors.New("username/email/password salah")
	// jwt error
	ErrJwtInvalidToken = errors.New("token tidak valid")
)
