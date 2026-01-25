package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/repository"
	"context"

	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUserProfile(ctx context.Context, userID uint) (*dto.UserResponse, error)
	UpdateUserProfile(ctx context.Context, userID uint, req dto.UserUpdateRequest) error
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

func (u *userUsecaseImpl) UpdateUserProfile(ctx context.Context,
	userID uint,
	req dto.UserUpdateRequest,
) error {

	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if user == nil {
		return gorm.ErrRecordNotFound
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Username != nil {
		user.Username = *req.Username
	}

	return u.userRepository.Update(ctx, user)
}

func (u *userUsecaseImpl) GetUserProfile(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:        user.PublicID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
