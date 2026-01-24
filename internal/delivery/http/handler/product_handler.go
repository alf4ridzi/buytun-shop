package handler

import (
	"buytun-backend/internal/usecase"

	"github.com/labstack/echo/v5"
)

type ProductHandler struct {
	uc usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) NewProduct(c *echo.Context) error {
	return nil
}
