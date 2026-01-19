package handler

import "buytun-backend/internal/usecase"

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: userUsecase}
}

func (h *UserHandler) Login() {

}
