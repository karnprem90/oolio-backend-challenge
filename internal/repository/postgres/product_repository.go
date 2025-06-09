package postgres

import (
	"context"
	"oolio-backend-challenge/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Get(ctx context.Context, id string) (domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRowxContext(ctx, "SELECT id, name, price, category FROM products WHERE id = $1", id).
		StructScan(&product)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return domain.Product{}, domain.ErrProductNotFound
		}
		return domain.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.SelectContext(ctx, &products, "SELECT id, name, price, category FROM products"); err != nil {
		return nil, err
	}
	return products, nil
}
