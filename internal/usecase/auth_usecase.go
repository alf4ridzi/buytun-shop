package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"buytun-backend/internal/utils"
	"buytun-backend/internal/utils/tokenutil"
	"context"
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Refresh(ctx context.Context, req dto.RefreshRequest) (*dto.AccessTokenResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthTokenResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) AuthUsecase {
	return &authUsecaseImpl{userRepository: userRepository}
}

func (u *authUsecaseImpl) Refresh(ctx context.Context, req dto.RefreshRequest) (
	*dto.AccessTokenResponse,
	error) {

	claims, err := tokenutil.ParseRefreshToken(req.Refresh)
	if err != nil {
		return nil, err
	}

	userID64, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return nil, err
	}

	userID := uint(userID64)

	user, err := u.userRepository.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	userIDString := strconv.Itoa(int(user.ID))

	tokenAccess, err := tokenutil.CreateUserAccessToken(userIDString, user.Name)
	if err != nil {
		return nil, err
	}

	resp := &dto.AccessTokenResponse{
		Access: tokenAccess,
	}

	return resp, nil
}

func (u *authUsecaseImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthTokenResponse, error) {
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

	userID := strconv.Itoa(int(user.ID))

	accessToken, err := tokenutil.CreateUserAccessToken(userID, user.Name)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokenutil.CreateUserRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	resp := &dto.AuthTokenResponse{
		Access:  accessToken,
		Refresh: refreshToken,
	}

	return resp, nil
}

func (u *authUsecaseImpl) Register(ctx context.Context, req dto.RegisterRequest) error {
	emailExist, err := u.userRepository.IsExistByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if !emailExist {
		return domain.ErrEmailAlreadyExist
	}

	usernameExist, err := u.userRepository.IsExistByUsername(ctx, req.Username)
	if err != nil {
		return err
	}

	if !usernameExist {
		return domain.ErrUsernameAlreadyExist
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	return u.userRepository.Create(ctx, user)
}
