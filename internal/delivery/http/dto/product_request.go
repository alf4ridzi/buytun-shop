package dto

type NewProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint64 `json:"price" validate:"required"`
	Stock       uint64 `json:"stock" validate:"required"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *uint64 `json:"price"`
	Stock       *uint64 `json:"stock"`
}
