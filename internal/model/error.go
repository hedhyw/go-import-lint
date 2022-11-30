package model

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"
)

// Error is a constant like error.
type Error string

func (err Error) Error() string {
	return string(err)
}

// errorsSet stores many errors.
type errorsSet struct {
	errs []error
}

// Error implements error interface.
func (err errorsSet) Error() string {
	errValues := make([]string, 0, len(err.errs))

	for _, e := range err.errs {
		errValues = append(errValues, e.Error())
	}

	return strings.Join(errValues, ", ")
}

// NewErrorSet creates error that handles many errors. It skips all nil
// errors.
func NewErrorSet(errs ...error) (err error) {
	actualErrs := make([]error, 0, len(errs))

	for _, err = range errs {
		if err == nil {
			continue
		}

		actualErrs = append(actualErrs, err)
	}

	switch len(actualErrs) {
	case 0:
		return nil
	case 1:
		return actualErrs[0]
	default:
		return errorsSet{
			errs: actualErrs,
		}
	}
}

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
	value, err := strconv.Unquote(n.Value)
	if err != nil {
		value = n.Value
	}

	return importOrderError{
		Position: n.Position,
		Reason:   reason,
		Value:    value,
	}
}
