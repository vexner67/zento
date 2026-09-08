package order

import "errors"

type Quantity int32

var ErrInvalidQuantity = errors.New("quantity must be greater than zero")

func NewQuantity(value int32) (Quantity, error) {
	if value <= 0 {
		return 0, ErrInvalidQuantity
	}

	return Quantity(value), nil
}

func (q Quantity) Int32() int32 {
	return int32(q)
}

func (q Quantity) Int64() int64 {
	return int64(q)
}
