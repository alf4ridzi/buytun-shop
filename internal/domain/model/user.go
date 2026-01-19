package model

import (
	"buytun-backend/internal/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"type:varchar(100);uniqueIndex:uq_email_user;not null"`
	Username  string `gorm:"type:varchar(100);uniqueIndex:uq_username_user; not null"`
	Password  string `gorm:"type:varchar(255); not null"`
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
