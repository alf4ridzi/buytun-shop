package repository

import "gorm.io/gorm"

type ProductRepository interface {
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}
