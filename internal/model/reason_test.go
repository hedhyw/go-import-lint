package model_test

import (
	"testing"

	"github.com/hedhyw/go-import-lint/internal/model"
)

func TestReasonFromError(t *testing.T) {
	t.Run("import_order_error", func(t *testing.T) {
		t.Parallel()

		const expReason = model.ReasonExtraLine
		err := model.NewImportOrderError(model.Node{}, expReason)

		gotReason := model.ReasonFromError(err)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		const expReason = model.ReasonUnknown

		gotReason := model.ReasonFromError(nil)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})

	t.Run("other error", func(t *testing.T) {
		t.Parallel()

		const expReason = model.ReasonUnknown
		const err model.Error = "test error"

		gotReason := model.ReasonFromError(err)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})
}
