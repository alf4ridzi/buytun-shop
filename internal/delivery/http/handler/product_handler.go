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

func (h *ProductHandler) GetProductByUser(c *echo.Context) error {
	userPublicID := c.Param("id")

	return response.Success(
		c,
		"ok",
		userPublicID,
	)
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

	userIDVal := c.Get("user_id")

	if userIDVal == nil {
		return response.Error(
			c,
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	resp, err := h.uc.AddNewProduct(
		c.Request().Context(),
		userID,
		req,
	)

	if err != nil {
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	return response.Success(
		c,
		"berhasil membuat product baru",
		resp,
	)
}
