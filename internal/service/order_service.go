package service

import (
	"context"

	"oolio-backend-challenge/internal/domain"

	"github.com/google/uuid"
)

type OrderService struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderService(
	orderRepo domain.OrderRepository,
	productRepo domain.ProductRepository,
) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) PlaceOrder(ctx context.Context, items []domain.OrderItem) (*domain.Order, error) {
	if len(items) == 0 {
		return nil, domain.ErrInvalidInput
	}

	products := make([]domain.Product, 0, len(items))
	for _, item := range items {
		if item.Quantity <= 0 {
			return nil, domain.ErrInvalidQuantity
		}

		product, err := s.productRepo.Get(ctx, item.ProductID)
		if err != nil {
			return nil, domain.ErrProductNotFound
		}
		products = append(products, product)
	}

	order := &domain.Order{
		ID:       uuid.New().String(),
		Items:    items,
		Products: products,
	}

	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	return s.orderRepo.Get(ctx, id)
}

func (s *OrderService) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return s.orderRepo.List(ctx)
}
