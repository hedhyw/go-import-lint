// Package model describes basic linter entities.
package model

import (
	"go/ast"
	"go/token"
	"strings"
)

// SuffixRecursive is a path suffix which enables recursive directories
// discovering.
const SuffixRecursive = "..."

// Node describes a unit from the program.
type Node struct {
	Kind     NodeKind
	Position token.Position
	Value    string
	// Offset out of comments.
	Offset int
}

// Index returns a line number subtracting offset.
func (n Node) Index() int {
	return n.Position.Line - n.Offset
}

// NewCommentNode creates new Node for comment entity.
func NewCommentNode(
	fset *token.FileSet,
	comment *ast.Comment,
) (n Node) {
	var position = fset.Position(comment.Pos())
	return Node{
		Kind:     KindComment,
		Value:    comment.Text,
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
	var kind NodeKind
	switch {
	case spec == nil, spec.Path == nil:
		return Node{
			Kind: KindUnknown,
		}
	case spec.Name != nil && spec.Name.Name == "_":
		kind = KindImportUnused
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

// NodeKind describes a type of the node.
type NodeKind uint8

// Possible node Kinds.
const (
	// KindUnknown is a undeterminated kind.
	KindUnknown NodeKind = iota
	// KindImportSTD is a standart library import kind.
	KindImportSTD
	// KindImportInternal is a current package import kind.
	KindImportInternal
	// KindImportVendor is a external package import kind.
	KindImportVendor
	// KindComment is a regular comment kind.
	KindComment
	// KindImportUnused is an unused package. For example SQL driver.
	KindImportUnused
)
