package model_test

import (
	"go/token"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/model"
)

func TestNewImportOrderError(t *testing.T) {
	const expReason = model.ReasonExtraLine
	var err = model.NewImportOrderError(model.Node{
		Kind:     model.KindImportInternal,
		Offset:   0,
		Position: token.Position{},
		Value:    "go/token",
	}, expReason)

	var gotReason = model.ReasonFromError(err)
	if expReason != gotReason {
		t.Fatalf("reason: got %s, exp %s", gotReason, expReason)
	}
}
