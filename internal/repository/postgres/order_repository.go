package postgres

import (
	"context"
	"encoding/json"
	"oolio-backend-challenge/internal/domain"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *domain.Order) error {
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx,
		"INSERT INTO orders (id, items, products) VALUES ($1, $2, $3)",
		order.ID, itemsJSON, productsJSON,
	)
	return err
}

func (r *OrderRepository) Get(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	var itemsJSON, productsJSON []byte

	err := r.db.QueryRowxContext(ctx,
		"SELECT id, items, products FROM orders WHERE id = $1",
		id,
	).Scan(&order.ID, &itemsJSON, &productsJSON)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	if err := json.Unmarshal(itemsJSON, &order.Items); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(productsJSON, &order.Products); err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) List(ctx context.Context) ([]domain.Order, error) {
	rows, err := r.db.QueryxContext(ctx, "SELECT id, items, products FROM orders ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var itemsJSON, productsJSON []byte

		if err := rows.Scan(&order.ID, &itemsJSON, &productsJSON); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(itemsJSON, &order.Items); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(productsJSON, &order.Products); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
