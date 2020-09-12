package model

import (
	"fmt"
	"go/token"
	"strconv"
)

type importOrderError struct {
	Position token.Position
	Value    string
	Reason   Reason
}

func (err importOrderError) Error() string {
	return fmt.Sprintf("%s: %s: %s", err.Position, err.Value, err.Reason)
}

func NewImportOrderError(el ImportElem, reason Reason) error {
	var value, err = strconv.Unquote(el.Value)
	if err != nil {
		value = el.Value
	}

	return importOrderError{
		Position: el.Position,
		Reason:   reason,
		Value:    value,
	}
}
