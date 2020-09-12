package model

import (
	"go/ast"
	"go/token"
	"strings"
)

// ImportElem
type ImportElem struct {
	Kind     importKind
	Position token.Position
	Spec     *ast.ImportSpec
	Value    string
}

// NewImportElem creates ImportElem.
func NewImportElem(
	fset *token.FileSet,
	spec *ast.ImportSpec,
	pkg string,
) (el ImportElem) {
	var kind importKind
	switch {
	case spec == nil, spec.Path == nil:
		return ImportElem{
			Kind: kindImportUnknown,
		}
	case !strings.Contains(spec.Path.Value, "."):
		kind = kindImportSTD
	case strings.Contains(spec.Path.Value, pkg):
		kind = KindImportInternal
	default:
		kind = kindImportVendor
	}

	var position = fset.Position(spec.Pos())

	return ImportElem{
		Kind:     kind,
		Position: position,
		Spec:     spec,
		Value:    spec.Path.Value,
	}
}

type importKind uint8

const (
	kindImportUnknown importKind = iota
	kindImportSTD
	KindImportInternal
	kindImportVendor
)
