package main

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/delivery/http/handler"
	"buytun-backend/internal/delivery/http/routes"
	"buytun-backend/internal/delivery/http/validator"
	"buytun-backend/internal/infrastructure/postgresql"
	"buytun-backend/internal/repository"
	"buytun-backend/internal/usecase"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func mustInitDB() *gorm.DB {
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

	return db
}

func closeDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}

func newEcho() *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()

	return e
}

func registerRoutes(e *echo.Echo, db *gorm.DB) {
	// user repo
	userRepo := repository.NewUserRepository(db)
	// product repo
	productRepo := repository.NewProductRepository(db)

	// auth
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)
	authRoute := routes.NewAuthRoute(authHandler)
	// product
	productUsecase := usecase.NewProductUsecase(productRepo, userRepo)
	productHandler := handler.NewProductHandler(productUsecase)
	productRoute := routes.NewProductRoute(productHandler)
	// user
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)
	userRoute := routes.NewUserRoute(userHandler, productHandler)

	route := routes.Routes{
		AuthRoute:    authRoute,
		UserRoute:    userRoute,
		ProductRoute: productRoute,
	}

	route.Register(e)
}

func startServer(e *echo.Echo, port int) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: e,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Error("failed to start server", "error", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		e.Logger.Error("failed to shutdown server", "error", err)
	}
}

func main() {
	cfg := config.LoadEnv()
	log.Printf("Port : %d", cfg.AppPort)

	db := mustInitDB()
	defer closeDB(db)

	e := newEcho()

	registerRoutes(e, db)
	startServer(e, cfg.AppPort)
}
