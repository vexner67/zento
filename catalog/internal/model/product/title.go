package product

import (
	"strings"
	"unicode/utf8"

	"github.com/vexner67/zento/catalog/internal/model/errors"
)

const (
	minTitleLength = 3
	maxTitleLength = 200
)

type Title string

func NewTitle(value string) (Title, error) {
	cleaned := strings.TrimSpace(value)
	length := utf8.RuneCountInString(cleaned)

	if length < minTitleLength {
		return "", errors.ErrTitleTooShort
	}

	if length > maxTitleLength {
		return "", errors.ErrTitleTooLong
	}

	return Title(cleaned), nil
}
