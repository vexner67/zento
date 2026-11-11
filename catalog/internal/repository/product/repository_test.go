package product

import (
	"testing"
	"uuid"

	"github.com/stretchr/testify/require"
	"github.com/vexner67/zento/catalog/internal/model/errors"
	"github.com/vexner67/zento/catalog/internal/model/product"
)

func createProduct() (product.Product, error) {
	title, err := product.NewTitle("MacBook M4")
	if err != nil {
		return product.Product{}, err
	}

	description, err := product.NewDescription("Очень мощный ноутбук")
	if err != nil {
		return product.Product{}, err
	}

	price, err := product.NewPrice(9_000_000)
	if err != nil {
		return product.Product{}, err
	}

	quantity, err := product.NewQuantity(7)
	if err != nil {
		return product.Product{}, err
	}

	return product.NewProduct(product.Params{
		SellerID:    uuid.NewV7(),
		CategoryID:  uuid.NewV7(),
		Title:       title,
		Description: description,
		ImageIDs:    []uuid.UUID{uuid.NewV7(), uuid.NewV7(), uuid.NewV7()},
		Price:       price,
		Quantity:    quantity,
	})
}

func TestSave(t *testing.T) {
	t.Parallel()

	t.Run("saves product", func(t *testing.T) {
		repo := NewRepository()

		p, err := createProduct()
		require.NoError(t, err)

		require.NoError(t, repo.Save(p))

		saved, err := repo.GetByID(p.ID())
		require.NoError(t, err)

		require.Equal(t, p, saved)
	})

	t.Run("allows only one concurrent save for the same ID", func(t *testing.T) {
		repo := NewRepository()

		p, err := createProduct()
		require.NoError(t, err)

		start := make(chan struct{})
		results := make(chan error, 2)

		for range 2 {
			go func() {
				<-start
				results <- repo.Save(p)
			}()
		}
		close(start)

		successes := 0
		for range 2 {
			if err = <-results; err == nil {
				successes++
			} else {
				require.ErrorIs(t, err, errors.ErrProductAlreadyExists)
			}
		}
		require.Equal(t, 1, successes)
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("replaces existing product", func(t *testing.T) {
		repo := NewRepository()

		p, err := createProduct()
		require.NoError(t, err)

		repo.storage[p.ID()] = product.Product{}

		require.NoError(t, repo.Update(p))

		updated, err := repo.GetByID(p.ID())
		require.NoError(t, err)

		require.Equal(t, p, updated)
	})

	t.Run("returns error for missing product", func(t *testing.T) {
		repo := NewRepository()

		p, err := createProduct()
		require.NoError(t, err)

		require.ErrorIs(t, repo.Update(p), errors.ErrProductNotFound)

		_, err = repo.GetByID(p.ID())
		require.ErrorIs(t, err, errors.ErrProductNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run("removes existing product", func(t *testing.T) {
		repo := NewRepository()

		p, err := createProduct()
		require.NoError(t, err)

		require.NoError(t, repo.Save(p))

		require.NoError(t, repo.Delete(p.ID()))

		_, err = repo.GetByID(p.ID())
		require.ErrorIs(t, err, errors.ErrProductNotFound)
	})

	t.Run("returns error for missing product", func(t *testing.T) {
		repo := NewRepository()

		require.ErrorIs(t, repo.Delete(uuid.NewV7()), errors.ErrProductNotFound)
	})
}
