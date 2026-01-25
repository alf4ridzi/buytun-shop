package model

import (
	"buytun-backend/internal/utils/slugutil"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Slug        string    `gorm:"uniqueIndex:idx_product_slug;not null"`
	Price       uint64    `gorm:"not null"`
	Stock       uint64    `gorm:"not null"`

	UserID uint `gorm:"index"`
	User   User `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.Slug = slugutil.GenerateProductSlug(p.Name)

	return nil
}
