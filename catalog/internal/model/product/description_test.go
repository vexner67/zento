package product

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

func TestNewDescription(t *testing.T) {
	t.Parallel()

	maxDescription := strings.Repeat("я", 5000)
	testCases := []struct {
		name    string
		value   string
		want    Description
		wantErr error
	}{
		{name: "empty"},
		{name: "blank", value: " \t\n "},
		{name: "trimmed", value: " \tКружка из керамики\n ", want: "Кружка из керамики"},
		{name: "maximum length after trimming", value: " " + maxDescription + " ", want: Description(maxDescription)},
		{name: "above maximum length", value: maxDescription + "я", wantErr: errors.ErrDescriptionTooLong},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewDescription(tc.value)
			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, got)
		})
	}
}
