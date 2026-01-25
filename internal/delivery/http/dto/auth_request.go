package dto

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh" validate:"required"`
}
