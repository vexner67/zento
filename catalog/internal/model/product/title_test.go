package product

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

func TestNewTitle(t *testing.T) {
	t.Parallel()

	maxTitle := strings.Repeat("я", 200)
	testCases := []struct {
		name    string
		value   string
		want    Title
		wantErr error
	}{
		{name: "minimum length after trimming", value: " \tКот\n ", want: "Кот"},
		{name: "maximum length after trimming", value: " " + maxTitle + " ", want: Title(maxTitle)},
		{name: "blank", value: " \t\n ", wantErr: errors.ErrTitleTooShort},
		{name: "below minimum length", value: "ёж", wantErr: errors.ErrTitleTooShort},
		{name: "above maximum length", value: maxTitle + "я", wantErr: errors.ErrTitleTooLong},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			title, err := NewTitle(tc.value)
			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, title)
		})
	}
}

func TestTitleString(t *testing.T) {
	t.Parallel()

	require.Equal(t, "Керамическая кружка", Title("Керамическая кружка").String())
}
