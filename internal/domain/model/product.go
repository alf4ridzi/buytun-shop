package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID   uint `gorm:"primaryKey"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
