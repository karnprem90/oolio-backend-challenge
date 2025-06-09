package service

import (
	"context"

	"oolio-backend-challenge/internal/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (domain.Product, error) {
	return s.repo.Get(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]domain.Product, error) {
	return s.repo.List(ctx)
}
