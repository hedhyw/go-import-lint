package model_test

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/model"
)

const pkg = "github.com/hedhyw/go-import-lint"

func TestNewCommentNode(t *testing.T) {
	gotNode := model.NewCommentNode(token.NewFileSet(), &ast.Comment{
		Slash: token.NoPos,
		Text:  "test",
	})
	expNode := model.Node{
		Kind:   model.KindComment,
		Offset: 0,
		Position: token.Position{
			Filename: gotNode.Position.Filename,
			Column:   0,
			Line:     0,
			Offset:   0,
		},
		Value: "test",
	}

	if !reflect.DeepEqual(expNode, gotNode) {
		t.Fatalf("kind: exp %+v, got %+v", expNode, gotNode)
	}
}

func TestNewImportNodeKind(t *testing.T) {
	t.Run("std_import", func(tt *testing.T) {
		testImportNodeKinds(tt, "fmt", model.KindImportSTD)
	})

	t.Run("internal_import", func(tt *testing.T) {
		testImportNodeKinds(tt, pkg+"/model", model.KindImportInternal)
	})

	t.Run("vendor_import", func(tt *testing.T) {
		testImportNodeKinds(tt, "github.com/hedhyw/jsoncjson", model.KindImportVendor)
	})
}

func testImportNodeKinds(
	t *testing.T,
	expValue string,
	expKind model.NodeKind,
) {
	expValue = strconv.Quote(expValue)

	n := model.NewImportNode(token.NewFileSet(), &ast.ImportSpec{
		Path: &ast.BasicLit{
			Value: expValue,
		},
	}, pkg)

	switch {
	case n.Kind != expKind:
		t.Fatalf("kind: exp %d, got %d", expKind, n.Kind)
	case n.Value != expValue:
		t.Fatalf("value: exp %s, got %s", expValue, n.Value)
	}
}

func TestUnknownImportKind(t *testing.T) {
	n := model.NewImportNode(token.NewFileSet(), &ast.ImportSpec{
		Path: nil,
	}, pkg)

	if n.Kind != model.KindUnknown {
		t.Fatalf("kind: exp %d, got %d", model.KindUnknown, n.Kind)
	}

	n = model.NewImportNode(token.NewFileSet(), nil, "")

	if n.Kind != model.KindUnknown {
		t.Fatalf("kind: exp %d, got %d", model.KindUnknown, n.Kind)
	}
}

func TestNodeIndex(t *testing.T) {
	const expIndex = 1

	n := model.Node{
		Offset: 1,
		Position: token.Position{
			Line: 2,
		},
	}

	if n.Index() != expIndex {
		t.Fatalf("index: exp %d, got %d", expIndex, n.Index())
	}
}
