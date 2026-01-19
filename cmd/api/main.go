package main

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/infrastructure/postgresql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

func main() {
	cfg := config.LoadEnv()
	log.Printf("Port : %d", cfg.AppPort)

	db, err := postgresql.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	if err := e.Start(fmt.Sprintf(":%d", cfg.AppPort)); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}
}
