package usecase

import (
	"buytun-backend/internal/delivery/http/dto"
	"buytun-backend/internal/domain/model"
	"buytun-backend/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ProductUsecase interface {
	GetProductByUserID(ctx context.Context, publicID uuid.UUID) ([]dto.ProductResponse, error)
	AddNewProduct(ctx context.Context, userID uint, req dto.NewProductRequest) (*dto.ProductResponse, error)
}

type productUsecaseImpl struct {
	repo     repository.ProductRepository
	userRepo repository.UserRepository
}

func NewProductUsecase(repo repository.ProductRepository, userRepo repository.UserRepository) ProductUsecase {
	return &productUsecaseImpl{repo: repo, userRepo: userRepo}
}

func (u *productUsecaseImpl) GetProductByUserID(ctx context.Context, publicID uuid.UUID) ([]dto.ProductResponse, error) {
	user, err := u.userRepo.FindByPublicID(ctx, publicID)
	if err != nil {
		return nil, err
	}

	products, err := u.repo.FindByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var resps []dto.ProductResponse

	for _, product := range products {

		resp := dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Slug:        product.Slug,
			Price:       product.Price,
			Stock:       product.Stock,
			Seller: &dto.UserResponse{
				ID:       product.User.PublicID,
				Name:     product.User.Name,
				Username: product.User.Username,
				Email:    product.User.Email,
			},
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}

		resps = append(resps, resp)
	}

	return resps, nil
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
