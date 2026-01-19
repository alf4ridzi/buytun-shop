package handler

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: authUsecase}
}

func (h *AuthHandler) Login() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}
}

func (h *AuthHandler) Register() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var req dto.RegisterRequest
		if err := c.Bind(&req); err != nil {
			return response.Error(
				c,
				http.StatusBadRequest,
				err.Error(),
				nil,
			)
		}

		return nil
	}
}
