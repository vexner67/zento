package product

import (
	"sync"
	"uuid"

	"github.com/vexner67/zento/catalog/internal/model/errors"
	"github.com/vexner67/zento/catalog/internal/model/product"
)

type repository struct {
	mu      sync.Mutex
	storage map[uuid.UUID]product.Product
}

func NewRepository() *repository {
	return &repository{
		storage: make(map[uuid.UUID]product.Product),
	}
}

func (r *repository) Save(p product.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[p.ID()]; exists {
		return errors.ErrProductAlreadyExists
	}

	r.storage[p.ID()] = p
	return nil
}

func (r *repository) GetByID(id uuid.UUID) (product.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, exists := r.storage[id]
	if !exists {
		return product.Product{}, errors.ErrProductNotFound
	}

	return p, nil
}

func (r *repository) Update(p product.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[p.ID()]; !exists {
		return errors.ErrProductNotFound
	}

	r.storage[p.ID()] = p
	return nil
}

func (r *repository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[id]; !exists {
		return errors.ErrProductNotFound
	}

	delete(r.storage, id)
	return nil
}
