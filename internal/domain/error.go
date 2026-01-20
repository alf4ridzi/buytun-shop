package domain

import "errors"

var (
	// auth error
	ErrInvalidAuth = errors.New("username/email/password salah")
	// user error
	ErrUserNotFound = errors.New("user tidak ditemukan")
	// jwt error
	ErrJwtInvalidToken = errors.New("token tidak valid")
)
