package product

import "github.com/vexner67/zento/catalog/internal/model/errors"

type Status string

const (
	StatusDraft    Status = "draft"
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
)

func NewStatus(value string) (Status, error) {
	switch Status(value) {
	case StatusDraft, StatusActive, StatusArchived:
		return Status(value), nil
	default:
		return "", errors.ErrInvalidProductStatus
	}
}
