package model

import (
	"go/ast"
	"go/token"
	"strings"
)

// Node describes a unit from the program.
type Node struct {
	Kind     nodeKind
	Position token.Position
	Value    string
	// Offset out of comments.
	Offset int
}

// Index returns a line number subtracting offset.
func (n Node) Index() int {
	return n.Position.Line - n.Offset
}

// NewImportNode creates new Node for comment entity.
func NewCommentNode(
	fset *token.FileSet,
	pos token.Pos,
) (n Node) {
	var position = fset.Position(pos)
	return Node{
		Kind:     KindComment,
		Value:    "",
		Position: position,
		Offset:   0,
	}
}

// NewImportNode creates new Node for import entity.
func NewImportNode(
	fset *token.FileSet,
	spec *ast.ImportSpec,
	pkg string,
) (n Node) {
	var kind nodeKind
	switch {
	case spec == nil, spec.Path == nil:
		return Node{
			Kind: KindImportUnknown,
		}
	case !strings.Contains(spec.Path.Value, "."):
		kind = KindImportSTD
	case strings.Contains(spec.Path.Value, pkg):
		kind = KindImportInternal
	default:
		kind = KindImportVendor
	}

	var position = fset.Position(spec.Pos())

	return Node{
		Kind:     kind,
		Value:    spec.Path.Value,
		Position: position,
		Offset:   0,
	}
}

type nodeKind uint8

const (
	KindImportUnknown nodeKind = iota
	KindImportSTD
	KindImportInternal
	KindImportVendor
	KindComment
)
