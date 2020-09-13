// Package walker contains the walker that scans directories for golang
// files.
package walker

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hedhyw/go-import-lint/internal/linter"
	"github.com/hedhyw/go-import-lint/internal/model"
)

const filesChSize = 1024
const extGo = ".go"

type walker struct {
	fset    *token.FileSet
	exclude []string

	foundFiles chan *ast.File
}

// Walker explores golang source files.
type Walker interface {
	Walk(path string) (err error)
	Files() <-chan *ast.File
	Close()
}

// NewWalker creates new Walker.
func NewWalker(fset *token.FileSet, exclude []string) (w Walker, err error) {
	for i, p := range exclude {
		exclude[i], err = filepath.Abs(p)
		if err != nil {
			return nil, fmt.Errorf("getting absolute path: %w", err)
		}
	}

	return walker{
		fset:    fset,
		exclude: exclude,

		foundFiles: make(chan *ast.File, filesChSize),
	}, nil
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
	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("getting absolute path: %w", err)
	}

	if strings.HasSuffix(path, model.SuffixRecursive) {
		err = filepath.Walk(
			strings.TrimSuffix(path, model.SuffixRecursive),
			w.filepathWalker,
		)

		if err != nil {
			return fmt.Errorf("walking: %w", err)
		}

		return nil
	}

	var info os.FileInfo
	info, err = os.Stat(path)
	switch {
	case err != nil:
		return fmt.Errorf("getting stat: %w", err)
	case info.IsDir():
		if err = w.handleDir(path); err != nil {
			return fmt.Errorf("handling dir: %w", err)
		}
	default:
		if err = w.handleFile(path); err != nil {
			return fmt.Errorf("handling file: %w", err)
		}
	}

	return nil
}

func (w walker) filepathWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		if err = w.shouldSkip(path); err != nil {
			return err
		}

		return nil
	}

	if err = w.handleFile(path); err != nil {
		return fmt.Errorf("handling file: %w", err)
	}

	return nil
}

func (w walker) handleDir(path string) (err error) {
	if err = w.shouldSkip(path); err != nil {
		return nil
	}

	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("reading dir: %w", err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		err = w.handleFile(filepath.Join(path, f.Name()))
		if err != nil {
			return fmt.Errorf("handling file: %w", err)
		}
	}

	return nil
}

func (w walker) shouldSkip(path string) (err error) {
	for _, excludePath := range w.exclude {
		if path == excludePath {
			return filepath.SkipDir
		}
	}

	return nil
}

func (w *walker) handleFile(filename string) (err error) {
	if err = w.shouldSkip(filename); err != nil {
		return nil
	}

	if !strings.HasSuffix(filename, extGo) {
		return nil
	}

	var f *ast.File
	f, err = parser.ParseFile(w.fset, filename, nil, linter.ParserMode)
	if err != nil {
		return fmt.Errorf("parsing file: %w", err)
	}

	w.foundFiles <- f

	return nil
}
