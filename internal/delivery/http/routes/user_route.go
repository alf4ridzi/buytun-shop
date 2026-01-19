package http

import (
	"buytun-backend/internal/delivery/http/handler"

	"github.com/labstack/echo/v5"
)

type UserRoute struct {
	Handler *handler.UserHandler
}

func NewUserRoute(handler *handler.UserHandler) *UserRoute {
	return &UserRoute{
		Handler: handler,
	}
}

func (r *UserRoute) Register(rg *echo.Group) {
	// users := rg.Group("/users")
	// {

	// }
}
