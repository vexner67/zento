package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vexner67/zento/order/internal/order"
)

type OrderRepository interface {
	Save(context.Context, *order.Order) error
}

type CreateOrderInput struct {
	CustomerID uuid.UUID
	Items      []order.Item
}

type Order struct {
	repo OrderRepository
}

func NewService(repo OrderRepository) *Order {
	return &Order{
		repo: repo,
	}
}

func (s *Order) Create(ctx context.Context, input CreateOrderInput) (*order.Order, error) {
	o, err := order.New(input.CustomerID, input.Items)
	if err != nil {
		return nil, err
	}

	if err = s.repo.Save(ctx, o); err != nil {
		return nil, err
	}

	return o, nil
}
