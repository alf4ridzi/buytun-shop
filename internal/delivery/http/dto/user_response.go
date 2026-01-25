package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
