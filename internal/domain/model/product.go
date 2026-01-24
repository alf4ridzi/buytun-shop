package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	Slug        string `gorm:"uniqueIndex:idx_product_slug;not null"`
	Price       uint64 `gorm:"not null"`

	UserID uint `gorm:"index"`
	User   User `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
