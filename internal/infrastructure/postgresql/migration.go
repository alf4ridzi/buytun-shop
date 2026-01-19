package postgresql

import (
	"buytun-backend/internal/domain/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
