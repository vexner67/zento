package product

import (
	"testing"
	"time"
	"uuid"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	t.Parallel()

	testCases := []Params{
		{
			SellerID:    uuid.NewV4(),
			CategoryID:  uuid.NewV4(),
			Title:       "Керамическая кружка",
			Description: "Кружка объёмом 350 мл",
			ImageIDs:    []uuid.UUID{uuid.NewV4(), uuid.NewV4()},
			Price:       129_900,
			Quantity:    5,
		},
		{
			SellerID:    uuid.NewV4(),
			CategoryID:  uuid.NewV4(),
			Title:       "Настольная лампа",
			Description: "Лампа с регулировкой яркости",
			ImageIDs:    nil,
			Price:       349_000,
			Quantity:    0,
		},
		{
			SellerID:    uuid.NewV4(),
			CategoryID:  uuid.NewV4(),
			Title:       "Блокнот",
			Description: "",
			ImageIDs:    []uuid.UUID{uuid.NewV4()},
			Price:       19_900,
			Quantity:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title.String(), func(t *testing.T) {
			t.Parallel()

			before := time.Now()
			p, err := NewProduct(tc)
			after := time.Now()

			require.NoError(t, err)
			require.NotNil(t, p)

			require.Equal(t, byte(7), p.id[6]>>4)
			require.Equal(t, tc.SellerID, p.sellerID)
			require.Equal(t, tc.CategoryID, p.categoryID)
			require.Equal(t, tc.Title, p.title)
			require.Equal(t, tc.Description, p.description)
			require.Equal(t, tc.ImageIDs, p.imageIDs)
			require.Equal(t, tc.Price, p.price)
			require.Equal(t, tc.Quantity, p.quantity)
			require.Equal(t, StatusDraft, p.status)

			require.False(t, p.createdAt.Before(before))
			require.False(t, p.createdAt.After(after))
			require.True(t, p.updatedAt.Equal(p.createdAt))
		})
	}
}
