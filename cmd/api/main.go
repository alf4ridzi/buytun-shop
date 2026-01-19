package main

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/delivery/http/handler"
	http "buytun-backend/internal/delivery/http/routes"
	"buytun-backend/internal/delivery/http/validator"
	"buytun-backend/internal/infrastructure/postgresql"
	"buytun-backend/internal/repository"
	"buytun-backend/internal/usecase"
	"fmt"
	"log"

	"github.com/labstack/echo/v5"
)

func main() {
	cfg := config.LoadEnv()
	log.Printf("Port : %d", cfg.AppPort)

	db, err := postgresql.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := postgresql.Migration(db); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Validator = validator.New()

	// repo
	userRepo := repository.NewUserRepository(db)
	// auth
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)
	authRoute := http.NewAuthRoute(authHandler)

	route := http.Routes{
		AuthRoute: authRoute,
	}

	route.Register(e)

	if err := e.Start(fmt.Sprintf(":%d", cfg.AppPort)); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}
}
