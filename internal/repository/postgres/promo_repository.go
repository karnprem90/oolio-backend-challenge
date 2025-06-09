package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PromoRepository struct {
	db *sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{
		db: db,
	}
}

func (r *PromoRepository) ValidatePromoCode(ctx context.Context, code string) (bool, error) {
	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM promo_codes WHERE code = $1", code)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
