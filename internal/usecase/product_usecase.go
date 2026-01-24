package usecase

import "buytun-backend/internal/repository"

type ProductUsecase interface{}

type productUsecaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{repo: repo}
}
