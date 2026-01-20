package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"buytun-backend/internal/utils"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecaseImpl{userRepository: userRepository}
}

func (u *authUsecaseImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error) {
	var user *model.User
	var err error

	if strings.Contains(req.Identifier, "@") {
		user, err = u.userRepository.FindByEmail(ctx, req.Identifier)

	} else {
		user, err = u.userRepository.FindByUsername(ctx, req.Identifier)
	}

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, gorm.ErrRecordNotFound
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	if !utils.ValidatePasswordHash(hashed, user.Password) {
		return nil, errors.New("invalid username/email/password")
	}

	resp := &dto.UserResponse{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	return resp, nil
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
