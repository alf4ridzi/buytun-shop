package response

import "github.com/labstack/echo/v5"

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *echo.Context, message string, data any) error {
	return JSON(c, 200, true, message, data)
}

func Error(c *echo.Context, code int, message string) error {
	return JSON(c, code, false, message, nil)
}
