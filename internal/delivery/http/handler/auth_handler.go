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

func (h *AuthHandler) Login(c *echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	if err := c.Validate(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	resp, err := h.uc.Login(c.Request().Context(), req)
	if err != nil {
		return response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return response.Success(
		c,
		"ok",
		resp,
	)
}

func (h *AuthHandler) Register(c *echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	if err := c.Validate(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	err := h.uc.Register(c.Request().Context(), req)

	if err != nil {
		return response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return response.Success(
		c,
		"berhasil register",
		nil,
	)
}
