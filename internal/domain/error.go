package domain

import "errors"

var (
	ErrInvalidAuth = errors.New("username/email/password salah")
)
