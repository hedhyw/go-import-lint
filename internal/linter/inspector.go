package linter

import (
	"go/ast"
	"go/token"
	"sort"

	"github.com/hedhyw/go-import-lint/internal/model"
)

type inspector interface {
	Inspect() (errs []error)
}

type fileInspector struct {
	fset *token.FileSet
	file *ast.File

	pkg string
}

// Inspect file for import errors.
func (l fileInspector) Inspect() (errs []error) {
	var importElems []model.ImportElem
	ast.Inspect(l.file, func(n ast.Node) bool {
		if spec, ok := n.(*ast.ImportSpec); ok {
			var el = model.NewImportElem(l.fset, spec, l.pkg)
			importElems = append(importElems, el)
		}

		return true
	})

	return l.analyseImports(importElems)
}

func (l fileInspector) analyseImports(importElems []model.ImportElem) (errs []error) {
	sort.Slice(importElems, func(i, j int) bool {
		return importElems[i].Position.Line < importElems[j].Position.Line
	})

	var err error
	for i, cur := range importElems {
		if i == 0 {
			continue
		}

		var prev = importElems[i-1]
		if err = l.checkRules(cur, prev); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (l fileInspector) checkRules(cur model.ImportElem, prev model.ImportElem) error {
	var lineDiff = cur.Position.Line - prev.Position.Line
	if prev.Kind == cur.Kind {
		if lineDiff > 1 {
			return model.NewImportOrderError(cur, model.ReasonExtraLine)
		}

		return nil
	}

	switch lineDiff {
	case 1:
		return model.NewImportOrderError(cur, model.ReasonMissingLine)
	case 2:
		// OK. Go on.
	default:
		return model.NewImportOrderError(cur, model.ReasonTooManyLines)
	}

	return nil
}
