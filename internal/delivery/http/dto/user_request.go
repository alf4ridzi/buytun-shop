package dto

type UserUpdateRequest struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Username    *string `json:"username"`
	PhoneNumber *string `json:"phone_number"`
}
