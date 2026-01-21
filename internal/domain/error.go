package domain

import "errors"

var (
	// auth error
	ErrInvalidAuth          = errors.New("username/email/password salah")
	ErrEmailAlreadyExist    = errors.New("email sudah terdaftar")
	ErrUsernameAlreadyExist = errors.New("username sudah terdaftar")
	// user error
	ErrUserNotFound = errors.New("user tidak ditemukan")
	// jwt error
	ErrJwtInvalidToken = errors.New("token tidak valid")
)
