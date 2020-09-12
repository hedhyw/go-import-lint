package walker

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const filesChSize = 1024
const recursive = "..."

type walker struct {
	fset *token.FileSet

	foundFiles chan *ast.File
}

// Walker explores golang source files.
type Walker interface {
	Walk(path string) (err error)
	Files() <-chan *ast.File
	Close()
}

// NewWalker creates new Walker.
func NewWalker(fset *token.FileSet) Walker {
	return walker{
		fset: fset,

		foundFiles: make(chan *ast.File, filesChSize),
	}
}

// Files returns channel with explored files.
func (w walker) Files() <-chan *ast.File {
	return w.foundFiles
}

// Close destroys walker.
func (w walker) Close() {
	close(w.foundFiles)
}

// Walk starts golang source files exploring.
func (w walker) Walk(path string) (err error) {
	if strings.HasSuffix(path, recursive) {
		return filepath.Walk(
			strings.TrimSuffix(path, recursive),
			w.filepathWalker,
		)
	}

	return w.handleDir(path)
}

func (w walker) filepathWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		err = w.handleDir(path)
		if err != nil {
			return fmt.Errorf("handling dir: %w", err)
		}
	}

	return nil
}

func (w walker) handleDir(path string) (err error) {
	var pkgs map[string]*ast.Package
	pkgs, err = parser.ParseDir(w.fset, path, func(info os.FileInfo) bool {
		return true
	}, 0)

	if err != nil {
		return fmt.Errorf("parsing dir: %w", err)
	}

	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			w.foundFiles <- f
		}
	}

	return nil
}
