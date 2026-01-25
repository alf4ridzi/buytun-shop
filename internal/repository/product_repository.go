package repository

import (
	"buytun-backend/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *model.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}
