package response

import "github.com/labstack/echo/v5"

func JSON(c *echo.Context,
	statusCode int,
	status bool,
	message string,
	data any,
) error {
	return c.JSON(statusCode, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
