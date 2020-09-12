package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"

	"github.com/hedhyw/go-import-lint/internal/linter"
	"github.com/hedhyw/go-import-lint/internal/walker"
)

func main() {
	var (
		path = flag.String("path", "./...", "path to lint")
		pkg  = flag.String("pkg", "-", "module package")
	)
	flag.Parse()

	var fset = token.NewFileSet()

	var walker = walker.NewWalker(fset)

	go func() {
		defer walker.Close()
		var werr = walker.Walk(*path)
		if werr != nil {
			fmt.Printf("walking error: %s", werr)
			os.Exit(1)
		}
	}()

	var linterGotErr = make(chan bool)
	go func() {
		var gotErr bool
		defer func() { linterGotErr <- gotErr }()

		var linter = linter.NewLinter(*pkg)
		for f := range walker.Files() {
			var errs = linter.Lint(fset, f)
			for _, err := range errs {
				gotErr = true
				fmt.Println(err)
			}
		}
	}()

	if <-linterGotErr {
		os.Exit(1)
	}
}
