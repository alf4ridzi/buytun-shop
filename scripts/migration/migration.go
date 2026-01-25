package main

import (
	"buytun-backend/internal/config"
	"buytun-backend/internal/domain"
	"buytun-backend/internal/infrastructure/postgresql"
	"log"

	"gorm.io/gorm"
)

func migration(db *gorm.DB) error {
	return db.AutoMigrate(domain.MIGRATION...)
}

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

	// var user string

	// err = db.Raw("select current_user").Scan(&user).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(user)

	err = migration(db)
	if err != nil {
		log.Fatal(err)
	}
}
