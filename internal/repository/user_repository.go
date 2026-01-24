package repository

import (
	"buytun-backend/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Update(ctx context.Context, user *model.User) error
	IsExistByEmail(ctx context.Context, email string) (bool, error)
	IsExistByUsername(ctx context.Context, username string) (bool, error)
	FindByID(ctx context.Context, id uint) (*model.User, error)
	IsExistByID(ctx context.Context, id uint) (bool, error)
	FindByUsernameRaw(ctx context.Context, username string) (*model.User, error)
	FindByEmailRaw(ctx context.Context, email string) (*model.User, error)
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

func (r *userRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	return r.DB.WithContext(ctx).Updates(user).Error
}

func (r *userRepositoryImpl) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).Error

	return count > 0, err
}

func (r *userRepositoryImpl) IsExistByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&model.User{}).
		Where("username = ?", username).
		Count(&count).Error

	return count > 0, err
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	tx := r.DB.WithContext(ctx).
		First(&user, "id = ?", id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *userRepositoryImpl) IsExistByID(ctx context.Context, id uint) (bool, error) {
	var count int64
	tx := r.DB.WithContext(ctx).
		Where("id = ?", id).
		Count(&count)

	if tx.Error != nil {
		return false, tx.Error
	}

	return count > 0, nil
}

func (r *userRepositoryImpl) FindByUsernameRaw(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		Raw("SELECT * FROM users WHERE username = ?", username).
		First(&user).
		Error
	return &user, err
}

func (r *userRepositoryImpl) FindByEmailRaw(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		Raw("SELECT * FROM users WHERE email = ?", email).
		First(&user).
		Error
	return &user, err
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		First(&user, "email = ?", email).
		Error
	return &user, err
}

func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		First(&user, "username = ?", username).
		Error
	return &user, err
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}
