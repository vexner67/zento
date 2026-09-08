package order

import "errors"

type Money struct {
	currencyCode string
	units        int64
	nanos        int32
}

var (
	ErrInvalidCurrencyCode = errors.New("invalid currency code")
	ErrNegativeMoney       = errors.New("money cannot be negative")
	ErrInvalidNanos        = errors.New("nanos is out of range")
)

func NewMoney(currencyCode string, units int64, nanos int32) (Money, error) {
	if len(currencyCode) != 3 {
		return Money{}, ErrInvalidCurrencyCode
	}

	if units < 0 || nanos < 0 {
		return Money{}, ErrNegativeMoney
	}

	if nanos > 999_999_999 {
		return Money{}, ErrInvalidNanos
	}

	return Money{
		currencyCode: currencyCode,
		units:        units,
		nanos:        nanos,
	}, nil
}

func (m Money) CurrencyCode() string {
	return m.currencyCode
}

func (m Money) Units() int64 {
	return m.units
}

func (m Money) Nanos() int32 {
	return m.nanos
}

func (m Money) NanosInt64() int64 {
	return int64(m.nanos)
}
