// Package gomod helps to determinate package name from the go.mod file.
package gomod

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hedhyw/go-import-lint/internal/model"
)

const fileGoMod = "go.mod"

// Packager gets package info from the go.mod.
type Packager interface {
	// Package returns current package name from the current path or from
	// the one of given paths.
	Package(paths []string) (pkg string, ok bool)
}

type packager struct{}

// NewPackager creates new packager.
func NewPackager() packager {
	return packager{}
}

const errPackageUnknown model.Error = "cannot determinate package from the go.mod"

// Package returns current package name from the current path or from
// the one of given paths.
func (p packager) Package(paths []string) (pkg string, err error) {
	paths = append([]string{"."}, paths...)

	for _, p := range paths {
		p = strings.TrimSuffix(p, model.SuffixRecursive)
		p = filepath.Join(p, fileGoMod)

		_, err = os.Stat(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", fmt.Errorf("getting stat: %w", err)
		}

		return readPackageFromGoMod(p)
	}

	return "", errPackageUnknown
}

func readPackageFromGoMod(file string) (pkg string, err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}

	defer func() { err = model.NewErrorSet(err, f.Close()) }()

	return scanForModule(f)
}

func scanForModule(r io.Reader) (pkg string, err error) {
	const prefixModule = "module "

	s := bufio.NewScanner(r)
	for s.Scan() {
		data := s.Text()
		if !strings.HasPrefix(data, prefixModule) {
			continue
		}

		return strings.TrimPrefix(data, prefixModule), nil
	}

	err = s.Err()
	if err != nil {
		return "", fmt.Errorf("scanning: %w", err)
	}

	return "", errPackageUnknown
}
