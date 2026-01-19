package repository

import "gorm.io/gorm"

type UserRepository interface{}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}
