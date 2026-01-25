package handler

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ProductHandler struct {
	uc usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) NewProduct(c *echo.Context) error {
	var req dto.NewProductRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return nil
}
