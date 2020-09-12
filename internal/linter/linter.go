package linter

import (
	"go/ast"
	"go/token"
)

type linter struct {
	pkg string
}

// Linter analysis file for correct imports order.
type Linter interface {
	Lint(fset *token.FileSet, file *ast.File) (errs []error)
}

// NewLinter creates new Linter.
func NewLinter(pkg string) Linter {
	return linter{
		pkg: pkg,
	}
}

// Lint inspects file for import errors.
func (a linter) Lint(
	fset *token.FileSet,
	file *ast.File,
) (errs []error) {
	var insp inspector = fileInspector{
		fset: fset,
		file: file,

		pkg: a.pkg,
	}

	return insp.Inspect()
}
