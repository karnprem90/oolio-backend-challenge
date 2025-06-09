package domain

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")

	ErrOrderNotFound = errors.New("order not found")

	ErrInvalidPromoCode = errors.New("invalid promo code")

	ErrInvalidInput = errors.New("invalid input")

	ErrInvalidQuantity = errors.New("invalid quantity")
)
