package repository

import (
	"buytun-backend/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).First(&user, "username = ?", username).Error
	return &user, err
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}
