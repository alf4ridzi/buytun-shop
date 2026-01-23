package middleware

import (
	"buytun-backend/internal/domain"
	"buytun-backend/internal/helpers/response"
	"buytun-backend/internal/utils/tokenutil"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return response.Error(
				c,
				http.StatusUnauthorized,
				"header authorization tidak ada",
			)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Error(
				c,
				http.StatusUnauthorized,
				"format autentikasi salah",
			)
		}

		tokenStr := parts[1]

		parse, err := tokenutil.ParseAccessToken(tokenStr)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrJwtInvalidToken):
				return response.Error(
					c,
					http.StatusUnauthorized,
					"token tidak valid atau kadaluarsa",
				)
			default:
				log.Println(err)
				return response.Error(
					c,
					http.StatusInternalServerError,
					"internal server error",
				)
			}
		}

		val, err := strconv.ParseUint(parse.Subject, 10, 64)
		if err != nil {
			log.Println(err)
			return response.Error(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
		}

		userID := uint(val)

		c.Set("user_id", userID)
		c.Set("name", parse.Name)

		return next(c)
	}
}
