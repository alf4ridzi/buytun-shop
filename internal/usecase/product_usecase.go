package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"context"
)

type ProductUsecase interface{}

type productUsecaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{repo: repo}
}

func (u *productUsecaseImpl) AddNewProduct(ctx context.Context, userID uint, req dto.NewProductRequest) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UserID:      userID,
	}

}
