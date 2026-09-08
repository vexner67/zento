package order

import "github.com/google/uuid"

type Item struct {
	productID uuid.UUID
	quantity  Quantity
	unitPrice Money
}

func NewItem(productID uuid.UUID, quantity Quantity, unitPrice Money) Item {
	return Item{
		productID: productID,
		quantity:  quantity,
		unitPrice: unitPrice,
	}
}

func (i Item) ProductID() uuid.UUID {
	return i.productID
}

func (i Item) Quantity() Quantity {
	return i.quantity
}

func (i Item) UnitPrice() Money {
	return i.unitPrice
}
