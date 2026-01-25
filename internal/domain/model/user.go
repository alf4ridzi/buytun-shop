package model

import (
	"buytun-backend/internal/utils"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primaryKey"`
	PublicID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex:idx_public_id"`
	Name     string
	Email    string `gorm:"type:varchar(100);uniqueIndex:uq_email_user;not null"`
	Username string `gorm:"type:varchar(100);uniqueIndex:uq_username_user; not null"`
	Password string `gorm:"type:varchar(255); not null"`

	Product []Product `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Username = strings.TrimSpace(u.Username)
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashed
	return nil
}
