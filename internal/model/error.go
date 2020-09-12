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

// NewImportOrderError creates new error about invalid import order.
func NewImportOrderError(n Node, reason Reason) error {
	var value, err = strconv.Unquote(n.Value)
	if err != nil {
		value = n.Value
	}

	return importOrderError{
		Position: n.Position,
		Reason:   reason,
		Value:    value,
	}
}
