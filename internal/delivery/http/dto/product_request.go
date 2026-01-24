package dto

type NewProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
}
