package http

import "github.com/labstack/echo/v5"

type Routes struct {
	UserRoute *UserRoute
	AuthRoute *AuthRoute
}

func NewRoutes(
	userRoute *UserRoute,
	authRoute *AuthRoute) *Routes {
	return &Routes{
		UserRoute: userRoute,
		AuthRoute: authRoute,
	}
}

func (r *Routes) Register(router *echo.Echo) {
	api := router.Group("/api")

	r.AuthRoute.Register(api)
	r.UserRoute.Register(api)
}
