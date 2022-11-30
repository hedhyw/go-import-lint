package gomod

import (
	"errors"
	"strings"
	"testing"
)

const expPkg = "github.com/hedhyw/go-import-lint"

func TestGoModScanner(t *testing.T) {
	p := NewPackager()
	gotPkg, err := p.Package([]string{"../.."})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if gotPkg != expPkg {
		t.Fatalf("pkg: exp %s, got %s", expPkg, gotPkg)
	}
}

func TestGoModScannerGoModNotFound(t *testing.T) {
	p := NewPackager()
	_, err := p.Package(nil)
	if !errors.Is(err, errPackageUnknown) {
		t.Fatalf("err: exp %s, got: %s", errPackageUnknown, err)
	}
}

func TestGoModScanForModule(t *testing.T) {
	const content = `
module github.com/hedhyw/go-import-lint

go 1.15
	`

	gotPkg, err := scanForModule(strings.NewReader(content))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if gotPkg != expPkg {
		t.Fatalf("pkg: exp %s, got %s", expPkg, gotPkg)
	}
}

func TestGoModScanForModuleInvalid(t *testing.T) {
	t.Run("invalid", func(tt *testing.T) {
		testScanUnknown(tt, `
MODULE github.com/hedhyw/go-import-lint

go 1.15
		`)
	})

	t.Run("not found", func(tt *testing.T) {
		testScanUnknown(tt, `go 1.15`)
	})
}

func testScanUnknown(t *testing.T, content string) {
	_, err := scanForModule(strings.NewReader(content))
	if !errors.Is(err, errPackageUnknown) {
		t.Fatalf("err: exp %s, got: %s", errPackageUnknown, err)
	}
}
