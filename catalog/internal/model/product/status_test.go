package product

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

func TestNewStatus(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		value   string
		want    Status
		wantErr error
	}{
		{name: "draft", value: "draft", want: StatusDraft},
		{name: "active", value: "active", want: StatusActive},
		{name: "archived", value: "archived", want: StatusArchived},
		{name: "empty", wantErr: errors.ErrInvalidProductStatus},
		{name: "unknown", value: "published", wantErr: errors.ErrInvalidProductStatus},
		{name: "uppercase", value: "DRAFT", wantErr: errors.ErrInvalidProductStatus},
		{name: "with whitespace", value: " draft ", wantErr: errors.ErrInvalidProductStatus},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			status, err := NewStatus(tc.value)
			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, status)
		})
	}
}
