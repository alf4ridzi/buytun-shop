package http

import (
	"buytun-backend/internal/delivery/http/handler"

	"github.com/labstack/echo/v5"
)

type AuthRoute struct {
	Handler *handler.AuthHandler
}

func NewAuthRoute(handler *handler.AuthHandler) *AuthRoute {
	return &AuthRoute{
		Handler: handler,
	}
}

func (r *AuthRoute) Register(rg *echo.Group) {
	auth := rg.Group("/auth")
	auth.POST("/login", r.Handler.Login)
	auth.POST("/register", r.Handler.Register)
}
