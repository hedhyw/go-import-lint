package linter_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/linter"
	"github.com/hedhyw/go-import-lint/internal/model"
)

const testPkg = "github.com/hedhyw/go-import-lint"
const testFile = "file_test.go"

func TestValidImport(t *testing.T) {
	var program = `
	package linter

	import (
		"fmt"
	
		"github.com/hedhyw/go-import-lint/internal/model"

		"github.com/hedhyw/jsonscjson"
	)
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{})
}

func TestExtraLines(t *testing.T) {
	var program = `
	package linter

	import (
		"fmt"

		"errors"

		"github.com/hedhyw/go-import-lint/internal/model"

		"github.com/hedhyw/go-import-lint/internal"

		"github.com/hedhyw/jsonscjson"

		"github.com/hedhyw/jsonscjson/main"
	)
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{
		model.ReasonExtraLine: 3,
	})
}

func TestTooManyLines(t *testing.T) {
	var program = `
	package linter

	import (
		"fmt"
		"errors"


		"github.com/hedhyw/go-import-lint/internal"


		"github.com/hedhyw/jsonscjson"
	)
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{
		model.ReasonTooManyLines: 2,
	})
}

func TestMissingLines(t *testing.T) {
	var program = `
	package linter

	import (
		"fmt"
		"errors"
		"github.com/hedhyw/go-import-lint/internal"
		"github.com/hedhyw/jsonscjson"
	)
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{
		model.ReasonMissingLine: 2,
	})
}

func TestValidSingleLineImport(t *testing.T) {
	var program = `
	package linter

	import "fmt"
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{})
}

func TestManySingleLineImports(t *testing.T) {
	var program = `
	package linter

	import "fmt"

	import "errors"
	`

	var l = linter.NewLinter(testPkg)
	var errs = l.Lint(mustParseProgram(t, program))
	assertReasonErrs(t, errs, map[model.Reason]int{
		model.ReasonExtraLine: 1,
	})
}

func mustParseProgram(t *testing.T, program string) (fset *token.FileSet, file *ast.File) {
	t.Helper()

	fset = token.NewFileSet()

	var err error
	file, err = parser.ParseFile(fset, testFile, []byte(program), 0)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return fset, file
}

func assertReasonErrs(t *testing.T, got []error, exp map[model.Reason]int) {
	t.Helper()

	var reasons = make(map[model.Reason]int)

	for _, err := range got {
		var r = model.ReasonFromError(err)
		if r == model.ReasonUnknown {
			t.Fatalf("unknown reason: %s", r)
		}

		reasons[r]++
	}

	if !reflect.DeepEqual(reasons, exp) {
		t.Fatalf("exp: %+v got: %+v", exp, reasons)
	}
}
