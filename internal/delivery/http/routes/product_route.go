package routes

import (
	"buytun-backend/internal/delivery/http/handler"

	"github.com/labstack/echo/v5"
)

type ProductRoute struct {
	Handler *handler.ProductHandler
}

func NewProductRoute(handler *handler.ProductHandler) *ProductRoute {
	return &ProductRoute{
		Handler: handler,
	}
}

func (r *ProductRoute) Register(rg *echo.Group) {
	product := rg.Group("/products")
	product.POST("", r.Handler.NewProduct)
}
