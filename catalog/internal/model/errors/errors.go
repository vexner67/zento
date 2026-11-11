package errors

import "errors"

var (
	ErrTitleTooShort = errors.New("product title is too short")
	ErrTitleTooLong  = errors.New("product title is too long")

	ErrDescriptionTooLong = errors.New("product description is too long")

	ErrInvalidPrice = errors.New("price must be greater than zero")

	ErrNegativeQuantity = errors.New("quantity cannot be negative")

	ErrInvalidProductStatus    = errors.New("invalid product status")
	ErrInvalidStatusTransition = errors.New("invalid product status transition")
)
