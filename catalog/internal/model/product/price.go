package product

import (
	"github.com/vexner67/zento/catalog/internal/model/errors"
)

type Price int64

func NewPrice(kopecks int64) (Price, error) {
	if kopecks <= 0 {
		return 0, errors.ErrInvalidPrice
	}

	return Price(kopecks), nil
}
