package dto

import "time"

type ProductResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Price       uint64 `json:"price"`
	Stock       uint64 `json:"stock"`

	Seller *UserResponse `json:"seller,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
