package handler

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: userUsecase}
}

func (h *UserHandler) UpdateUser(c *echo.Context) error {
	var req dto.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
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
			"internal server error userid",
		)
	}

	err := h.uc.UpdateUserProfile(
		c.Request().Context(),
		userID,
		req,
	)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return response.Error(
				c,
				http.StatusNotFound,
				"user tidak ditemukan",
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
		"berhasil update user",
		nil,
	)

}

func (h *UserHandler) GetMe(c *echo.Context) error {
	val := c.Get("user_id")
	if val == nil {
		return response.Error(
			c,
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	userID, ok := val.(uint)
	if !ok {
		return response.Error(
			c,
			http.StatusInternalServerError,
			"internal server error",
		)
	}

	resp, err := h.uc.GetUserProfile(
		c.Request().Context(),
		userID,
	)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return response.Error(
				c,
				http.StatusNotFound,
				"user tidak ditemukan",
			)
		default:
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}
	}

	return response.Success(
		c,
		"berhasil mendapatkan user",
		resp,
	)
}
