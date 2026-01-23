package handler

import (
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: userUsecase}
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

}
