package linter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"

	"github.com/hedhyw/go-import-lint/internal/model"
)

// ParserMode is a mode that stops parsing after imports and includes
// comments.
const ParserMode = parser.ImportsOnly + parser.ParseComments

const (
	prefixGeneratedComment = "Code generated by"
	prefixNolintComment    = "nolint:go-import-lint"
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
	if shouldSkip(l.file.Comments) {
		return nil
	}

	var nodes []model.Node
	ast.Inspect(l.file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.ImportSpec:
			n := model.NewImportNode(l.fset, node, l.pkg)
			nodes = append(nodes, n)
		case *ast.Comment:
			n := model.NewCommentNode(l.fset, node)
			nodes = append(nodes, n)
		}

		return true
	})

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Index() < nodes[j].Index()
	})

	nodes = removeComments(nodes)
	return l.analyseImports(nodes)
}

func shouldSkip(comments []*ast.CommentGroup) (skip bool) {
	for _, c := range comments {
		if c == nil {
			continue
		}

		comment := strings.TrimSpace(c.Text())

		switch {
		case strings.HasPrefix(comment, prefixGeneratedComment):
			return true
		case strings.HasPrefix(comment, prefixNolintComment):
			return true
		}
	}

	return false
}

func removeComments(nodes []model.Node) (filteredNodes []model.Node) {
	filteredNodes = make([]model.Node, 0, len(nodes))

	var offset int
	for _, n := range nodes {
		switch n.Kind {
		case model.KindComment:
			offset++
		default:
			n.Offset = offset
			filteredNodes = append(filteredNodes, n)
		}
	}

	return filteredNodes
}

func (l fileInspector) analyseImports(nodes []model.Node) (errs []error) {
	var err error
	for i, cur := range nodes {
		if i == 0 {
			continue
		}

		prev := nodes[i-1]
		if err = l.checkRules(cur, prev); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (l *fileInspector) checkRules(cur model.Node, prev model.Node) error {
	lineDiff := cur.Index() - prev.Index()
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
		// Import "C" is exceptional, because it should be in separate block.
		if cur.Kind != model.KindImportC && prev.Kind != model.KindImportC {
			return model.NewImportOrderError(cur, model.ReasonTooManyLines)
		}
	}

	return nil
}
