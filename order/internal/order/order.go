package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id         uuid.UUID
	customerID uuid.UUID
	items      []Item
	status     Status
	total      Money
	createdAt  time.Time
	updatedAt  time.Time
}

var ErrInvalidCustomerID = errors.New("customer id must not be empty")

func New(customerID uuid.UUID, items []Item) (*Order, error) {
	if customerID == uuid.Nil {
		return nil, ErrInvalidCustomerID
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	total, err := calculateTotal(items)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Order{
		id:         id,
		customerID: customerID,
		items:      append([]Item(nil), items...),
		status:     Pending,
		total:      total,
		createdAt:  now,
		updatedAt:  now,
	}, nil
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) CustomerID() uuid.UUID {
	return o.customerID
}

func (o *Order) Items() []Item {
	return append([]Item(nil), o.items...)
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) Total() Money {
	return o.total
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

var (
	ErrOrderHasNoItems = errors.New("order must contain at least one item")
	ErrMixedCurrencies = errors.New("order items must use the same currency")
)

func calculateTotal(items []Item) (Money, error) {
	if len(items) == 0 {
		return Money{}, ErrOrderHasNoItems
	}

	currencyCode := items[0].unitPrice.currencyCode

	var units int64
	var nanos int64

	for _, item := range items {
		if item.unitPrice.currencyCode != currencyCode {
			return Money{}, ErrMixedCurrencies
		}

		quantity := item.quantity.Int64()

		units += quantity * item.unitPrice.units
		nanos += quantity * item.unitPrice.NanosInt64()
	}

	nanosPerUnit := int64(1_000_000_000)

	units += nanos / nanosPerUnit
	nanos %= nanosPerUnit

	return NewMoney(currencyCode, units, int32(nanos))
}
