package repository

import (
	"buytun-backend/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByUserID(ctx context.Context, userID uint) ([]model.Product, error)
	Create(ctx context.Context, product *model.Product) error
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (r *productRepositoryImpl) FindByUserID(ctx context.Context, userID uint) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.WithContext(ctx).
		Model(&model.Product{}).
		Where("user_id = ?", userID).
		Preload("User").
		Find(&products).Error

	return products, err
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *model.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}
