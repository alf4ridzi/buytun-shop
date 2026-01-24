package routes

import (
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	buytunmiddleware "buytun-backend/internal/delivery/http/middleware"
)

type Routes struct {
	UserRoute    *UserRoute
	AuthRoute    *AuthRoute
	ProductRoute *ProductRoute
}

func NewRoutes(
	userRoute *UserRoute,
	authRoute *AuthRoute,
	productRoute *ProductRoute) *Routes {
	return &Routes{
		UserRoute:    userRoute,
		AuthRoute:    authRoute,
		ProductRoute: productRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")
	api.Use(middleware.ContextTimeoutWithConfig(middleware.ContextTimeoutConfig{
		Timeout: 2 * time.Second,
	}))

	r.AuthRoute.Register(api)

	api.Use(buytunmiddleware.AuthMiddleware)

	r.UserRoute.Register(api)

	r.ProductRoute.Register(api)
}
