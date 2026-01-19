package handler

import (
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
