package handler

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProductHandler struct {
	uc usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) Update(c *echo.Context) error {
	return nil
}

func (h *ProductHandler) GetProductByUser(c *echo.Context) error {
	userPublicID := c.Param("id")

	if userPublicID == "" {
		return response.Error(
			c,
			http.StatusBadRequest,
			"user id tidak ada",
		)
	}

	userPublicUUID, err := uuid.Parse(userPublicID)
	if err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			"id tidak valid",
		)
	}

	products, err := h.uc.GetProductByUserID(
		c.Request().Context(),
		userPublicUUID,
	)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return response.Error(
				c,
				http.StatusNotFound,
				"product tidak ditemukan",
			)
		default:
			log.Println(err)
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}
	}

	return response.Success(
		c,
		"berhasil mendapatkan product",
		products,
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
