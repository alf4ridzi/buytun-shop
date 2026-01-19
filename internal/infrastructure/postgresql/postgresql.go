package postgresql

import (
	"buytun-backend/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.GetConfig().DBHost,
		config.GetConfig().DBPassword,
		config.GetConfig().DBName,
		config.GetConfig().DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
