package model_test

import (
	"testing"

	"github.com/hedhyw/go-import-lint/internal/model"
)

func TestReasonFromError(t *testing.T) {
	t.Run("import_order_error", func(tt *testing.T) {
		const expReason = model.ReasonExtraLine
		var err = model.NewImportOrderError(model.Node{}, expReason)

		var gotReason = model.ReasonFromError(err)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})

	t.Run("nil", func(tt *testing.T) {
		const expReason = model.ReasonUnknown

		var gotReason = model.ReasonFromError(nil)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})

	t.Run("other error", func(tt *testing.T) {
		const expReason = model.ReasonUnknown
		const err model.Error = "test error"

		var gotReason = model.ReasonFromError(err)
		if gotReason != expReason {
			t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
		}
	})
}
