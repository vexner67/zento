package product

import (
	"strings"
	"unicode/utf8"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

const maxDescriptionLength = 5000

type Description string

func NewDescription(value string) (Description, error) {
	cleaned := strings.TrimSpace(value)
	length := utf8.RuneCountInString(cleaned)

	if length > maxDescriptionLength {
		return "", errors.ErrDescriptionTooLong
	}

	return Description(cleaned), nil
}
