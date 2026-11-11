package product

import "github.com/vexner67/zento/catalog/internal/model/errors"

type Quantity int

func NewQuantity(value int) (Quantity, error) {
	if value < 0 {
		return 0, errors.ErrNegativeQuantity
	}

	return Quantity(value), nil
}
