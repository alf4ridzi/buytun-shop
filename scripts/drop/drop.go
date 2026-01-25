package main

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/domain"
	"buytun-backend/internal/infrastructure/postgresql"
	"log"
	"slices"

	"gorm.io/gorm"
)

func mustinitdb() (*gorm.DB, error) {
	db, err := postgresql.NewPostgresDB()
	return db, err
}

func closedb(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func main() {
	_ = config.LoadEnv()

	db, err := mustinitdb()
	if err != nil {
		log.Fatal(err)
	}

	defer closedb(db)

	reverse := domain.MIGRATION

	slices.Reverse(reverse)

	db.Migrator().DropTable(reverse...)
}
