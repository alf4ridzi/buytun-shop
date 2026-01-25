package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"context"
)

type ProductUsecase interface {
	AddNewProduct(ctx context.Context, userID uint, req dto.NewProductRequest) (*dto.ProductResponse, error)
}

type productUsecaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{repo: repo}
}

func (u *productUsecaseImpl) AddNewProduct(ctx context.Context, userID uint, req dto.NewProductRequest) (*dto.ProductResponse, error) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UserID:      userID,
	}

	err := u.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	resp := &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Slug:        product.Slug,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	return resp, nil
}
