package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vexner67/zento/order/internal/order"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

const insertOrderQuery = `
	INSERT INTO orders (
		id,
		customer_id,
		status,
		currency_code,
		total_units,
		total_nanos,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

const insertOrderItemQuery = `
	INSERT INTO order_items (
		order_id,
		product_id,
		quantity,
		unit_price_units,
		unit_price_nanos
	)
	VALUES ($1, $2, $3, $4, $5)
`

func (r *Repository) Save(ctx context.Context, o *order.Order) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(
		ctx,
		insertOrderQuery,
		o.ID(),
		o.CustomerID(),
		o.Status().String(),
		o.Total().CurrencyCode(),
		o.Total().Units(),
		o.Total().Nanos(),
		o.CreatedAt(),
		o.UpdatedAt(),
	)
	if err != nil {
		return fmt.Errorf("insert order: %w", err)
	}

	for _, item := range o.Items() {
		_, err = tx.Exec(
			ctx,
			insertOrderItemQuery,
			o.ID(),
			item.ProductID(),
			item.Quantity().Int32(),
			item.UnitPrice().Units(),
			item.UnitPrice().Nanos(),
		)
		if err != nil {
			return fmt.Errorf("insert order item: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit order transaction: %w", err)
	}

	return nil
}
