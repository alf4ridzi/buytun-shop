package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"context"
)

type AuthUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecaseImpl{userRepository: userRepository}
}

func (u *authUsecaseImpl) Register(ctx context.Context, req dto.RegisterRequest) error {
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	return u.userRepository.Create(ctx, user)
}
