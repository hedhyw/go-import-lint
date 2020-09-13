package gomod

import (
	"errors"
	"strings"
	"testing"
)

const expPkg = "github.com/hedhyw/go-import-lint"

func TestGoModScanner(t *testing.T) {
	var p = NewPackager()
	var gotPkg, err = p.Package([]string{"../.."})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if gotPkg != expPkg {
		t.Fatalf("pkg: exp %s, got %s", expPkg, gotPkg)
	}
}

func TestGoModScannerGoModNotFound(t *testing.T) {
	var p = NewPackager()
	var _, err = p.Package(nil)
	if !errors.Is(err, errPackageUnknown) {
		t.Fatalf("err: exp %s, got: %s", errPackageUnknown, err)
	}
}

func TestGoModScan(t *testing.T) {

	const content = `
module github.com/hedhyw/go-import-lint

go 1.14
	`

	var gotPkg, err = scanForModule(strings.NewReader(content))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if gotPkg != expPkg {
		t.Fatalf("pkg: exp %s, got %s", expPkg, gotPkg)
	}
}

func TestGoModScanInvalid(t *testing.T) {
	t.Run("invalid", func(tt *testing.T) {
		testScanUnknown(tt, `
MODULE github.com/hedhyw/go-import-lint

go 1.14
		`)
	})

	t.Run("not found", func(tt *testing.T) {
		testScanUnknown(tt, `go 1.14`)
	})
}

func testScanUnknown(t *testing.T, content string) {
	var _, err = scanForModule(strings.NewReader(content))
	if !errors.Is(err, errPackageUnknown) {
		t.Fatalf("err: exp %s, got: %s", errPackageUnknown, err)
	}
}
