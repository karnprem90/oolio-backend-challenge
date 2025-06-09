package domain

import "context"

type ProductRepository interface {
	Get(ctx context.Context, id string) (Product, error)

	List(ctx context.Context) ([]Product, error)
}

type OrderRepository interface {
	Create(ctx context.Context, order *Order) error

	Get(ctx context.Context, id string) (*Order, error)

	List(ctx context.Context) ([]Order, error)
}

type PromoRepository interface {
	ValidatePromoCode(ctx context.Context, code string) (bool, error)
}
