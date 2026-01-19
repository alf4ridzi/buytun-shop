package response

import "github.com/labstack/echo/v5"

type Response struct {
	c          *echo.Context
	StatusCode int
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
