package product

import (
	"time"
	"uuid"
)

type Product struct {
	id         uuid.UUID
	sellerID   uuid.UUID
	categoryID uuid.UUID

	title       Title
	description Description
	imageIDs    []uuid.UUID

	price    Price
	quantity Quantity
	status   Status

	updatedAt time.Time
	createdAt time.Time
}

type Params struct {
	SellerID    uuid.UUID
	CategoryID  uuid.UUID
	Title       Title
	Description Description
	ImageIDs    []uuid.UUID
	Price       Price
	Quantity    Quantity
}

func NewProduct(params Params) (*Product, error) {
	status, err := NewStatus("draft")
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Product{
		id:          uuid.NewV7(),
		sellerID:    params.SellerID,
		categoryID:  params.CategoryID,
		title:       params.Title,
		description: params.Description,
		imageIDs:    params.ImageIDs,
		price:       params.Price,
		quantity:    params.Quantity,
		status:      status,
		updatedAt:   now,
		createdAt:   now,
	}, nil
}
