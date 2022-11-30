package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"
	"strings"

	"github.com/hedhyw/go-import-lint/internal/gomod"
	"github.com/hedhyw/go-import-lint/internal/linter"
	"github.com/hedhyw/go-import-lint/internal/walker"
)

type stringsFlag struct {
	values    []string
	isInitial bool
}

func (f stringsFlag) String() string {
	return strings.Join([]string(f.values), ", ")
}

func (f *stringsFlag) Set(value string) error {
	if value == "" {
		return nil
	}

	if f.isInitial {
		f.values = make([]string, 0)
		f.isInitial = false
	}

	f.values = append(f.values, value)
	return nil
}

type flags struct {
	Package string
	Paths   stringsFlag
	Exclude stringsFlag
}

func newFlags() (f *flags) {
	f = new(flags)

	f.Paths = stringsFlag{[]string{"./..."}, true}
	f.Exclude = stringsFlag{[]string{"./vendor", "./.git"}, true}

	flag.StringVar(&f.Package, "pkg", "", "module package")
	flag.Var(&f.Paths, "path", "paths to lint")
	flag.Var(&f.Exclude, "exclude", "paths to exclude")

	return f
}

func main() {
	flagArgs := newFlags()
	flag.Parse()

	os.Exit(run(flagArgs))
}

func run(flagArgs *flags) (code int) {
	fset := token.NewFileSet()

	walker, err := walker.NewWalker(fset, flagArgs.Exclude.values)
	if err != nil {
		fmt.Printf("creating walker: %s\n", err)
		return 1
	}

	if flagArgs.Package == "" {
		p := gomod.NewPackager()
		flagArgs.Package, err = p.Package(flagArgs.Paths.values)
		if err != nil {
			fmt.Printf("getting package: %s\n", err)
			return 1
		}
	}

	walkerErr := make(chan error, 1)
	go func() {
		defer walker.Close()
		for _, p := range flagArgs.Paths.values {
			walkerErr <- walker.Walk(p)
		}
	}()

	linterGotErr := make(chan bool, 1)
	go func() {
		var gotErr bool
		defer func() { linterGotErr <- gotErr }()

		linter := linter.NewLinter(flagArgs.Package)
		for f := range walker.Files() {
			errs := linter.Lint(fset, f)
			for _, err := range errs {
				gotErr = true
				fmt.Println(err)
			}
		}
	}()

	if <-linterGotErr {
		return 1
	}

	werr := <-walkerErr
	if werr != nil {
		return 1
	}

	return 0
}
