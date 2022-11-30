package model_test

import (
	"go/token"
	"strings"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/model"
)

func TestConstError(t *testing.T) {
	const msg = "test"
	const err model.Error = msg

	if err.Error() != msg {
		t.Fatalf("err: got %s, exp %s", err, msg)
	}
}

func TestErrorSet(t *testing.T) {
	const (
		firstError  model.Error = "first test error"
		secondError model.Error = "second test error"
	)

	t.Run("two", func(tt *testing.T) {
		err := model.NewErrorSet(firstError, secondError)

		switch {
		case !strings.Contains(err.Error(), firstError.Error()):
			tt.Fatalf("err: got %s, exp include %s", err.Error(), firstError.Error())
		case !strings.Contains(err.Error(), secondError.Error()):
			tt.Fatalf("err: got %s, exp include %s", err.Error(), secondError.Error())
		}
	})

	t.Run("single", func(tt *testing.T) {
		err := model.NewErrorSet(firstError)

		if err != firstError {
			tt.Fatalf("err: got %s, exp %s", err, firstError)
		}
	})

	t.Run("single_nil", func(tt *testing.T) {
		err := model.NewErrorSet(firstError, nil)

		if err != firstError {
			tt.Fatalf("err: got %s, exp %s", err, firstError)
		}
	})

	t.Run("nil_single", func(tt *testing.T) {
		err := model.NewErrorSet(nil, secondError)

		if err != secondError {
			tt.Fatalf("err: got %s, exp %s", err, secondError)
		}
	})

	t.Run("empty", func(tt *testing.T) {
		err := model.NewErrorSet()

		if err != nil {
			tt.Fatalf("err: got %s, exp nil", err)
		}
	})

	t.Run("nil", func(tt *testing.T) {
		err := model.NewErrorSet(nil)

		if err != nil {
			tt.Fatalf("err: got %s, exp nil", err)
		}
	})
}

func TestNewImportOrderError(t *testing.T) {
	const expReason = model.ReasonExtraLine
	err := model.NewImportOrderError(model.Node{
		Kind:     model.KindImportInternal,
		Offset:   0,
		Position: token.Position{},
		Value:    "go/token",
	}, expReason)

	gotReason := model.ReasonFromError(err)
	if expReason != gotReason {
		t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
	}
}
