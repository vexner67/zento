package product

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

func TestNewQuantity(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		value   int
		want    Quantity
		wantErr error
	}{
		{name: "positive", value: 5, want: 5},
		{name: "zero", value: 0, want: 0},
		{name: "negative", value: -1, wantErr: errors.ErrNegativeQuantity},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			quantity, err := NewQuantity(tc.value)
			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, quantity)
		})
	}
}
