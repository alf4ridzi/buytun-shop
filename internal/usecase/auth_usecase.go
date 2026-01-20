package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"buytun-backend/internal/utils"
	"context"
	"strings"
)

type AuthUsecase interface {
	Login(ctx context.Context, req dto.LoginRequest) (*model.User, error)
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecaseImpl{userRepository: userRepository}
}

func (u *authUsecaseImpl) Login(ctx context.Context, req dto.LoginRequest) (*model.User, error) {
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
		return nil, domain.ErrInvalidAuth
	}

	if !utils.ValidatePasswordHash(user.Password, req.Password) {
		return nil, domain.ErrInvalidAuth
	}

	return user, nil
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
