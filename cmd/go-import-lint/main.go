package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/hedhyw/go-import-lint/internal/linter"
)

func main() {
	var (
		file = flag.String("file", "", "file to lint")
		pkg  = flag.String("pkg", "-", "file pkg")
	)
	flag.Parse()

	var fset = token.NewFileSet()
	var f, err = parser.ParseFile(fset, *file, nil, 0)
	if err != nil {
		fmt.Printf("cannot parse file: %s", err)
		os.Exit(1)
	}

	var linter = linter.NewLinter(*pkg)

	var errs = linter.Lint(fset, f)
	for _, err = range errs {
		fmt.Println(err)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}
