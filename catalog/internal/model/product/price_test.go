package product

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

func TestNewPrice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		kopecks int64
		want    Price
		wantErr error
	}{
		{name: "one kopeck", kopecks: 1, want: 1},
		{name: "maximum int64", kopecks: math.MaxInt64, want: math.MaxInt64},
		{name: "zero", kopecks: 0, wantErr: errors.ErrInvalidPrice},
		{name: "negative", kopecks: -1, wantErr: errors.ErrInvalidPrice},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			price, err := NewPrice(tc.kopecks)
			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, price)
		})
	}
}
