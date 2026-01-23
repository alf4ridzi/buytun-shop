package handler

import (
	"buytun-backend/internal/usecase"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: userUsecase}
}

func (h *UserHandler) GetMe(c *echo.Context) error {
	return nil
}
