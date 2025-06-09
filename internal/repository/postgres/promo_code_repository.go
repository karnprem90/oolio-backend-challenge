package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"oolio-backend-challenge/internal/domain"
)

type PromoCodeRepository struct {
	db *sql.DB
}

func NewPromoCodeRepository(db *sql.DB) *PromoCodeRepository {
	return &PromoCodeRepository{db: db}
}

func (r *PromoCodeRepository) ValidatePromoCode(ctx context.Context, code string) (bool, error) {
	if len(code) < 8 || len(code) > 10 {
		return false, domain.ErrInvalidPromoCode
	}

	query := `
		SELECT COUNT(DISTINCT file_name) 
		FROM promo_codes 
		WHERE code = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, strings.ToUpper(code)).Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return count >= 2, nil
}
