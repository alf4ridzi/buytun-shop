package usecase

import "buytun-backend/internal/repository"

type AuthUsecase interface{}

type authUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}
