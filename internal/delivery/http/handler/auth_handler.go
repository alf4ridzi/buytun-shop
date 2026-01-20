package handler

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: authUsecase}
}

func (h *AuthHandler) Refresh(c *echo.Context) error {
	var req dto.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			"failed to bind",
		)
	}

	if err := c.Validate(req); err != nil {
		return response.Error(
			c,
			http.StatusBadRequest,
			"bad request",
		)
	}

	return response.Success(
		c,
		"ok",
		nil,
	)
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
		switch {
		case errors.Is(err, domain.ErrInvalidAuth):
			return response.Error(
				c,
				http.StatusForbidden,
				err.Error(),
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
		"berhasil login",
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
