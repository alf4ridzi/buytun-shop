package usecase

import "buytun-backend/internal/repository"

type UserUsecase interface{}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}
