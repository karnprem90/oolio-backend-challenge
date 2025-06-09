package domain

import "context"

type ProductService interface {
	GetProduct(ctx context.Context, id string) (Product, error)

	ListProducts(ctx context.Context) ([]Product, error)
}

type OrderService interface {
	PlaceOrder(ctx context.Context, items []OrderItem) (*Order, error)

	GetOrder(ctx context.Context, id string) (*Order, error)

	ListOrders(ctx context.Context) ([]Order, error)
}

type PromoCodeService interface {
	ValidatePromoCode(ctx context.Context, code string) (bool, error)
}
