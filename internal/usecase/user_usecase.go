package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/repository"
	"context"
)

type UserUsecase interface {
	GetUserProfile(ctx context.Context, userID uint) (*dto.UserResponse, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

func (u *userUsecaseImpl) GetUserProfile(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
