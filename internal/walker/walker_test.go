package walker_test

import (
	"go/token"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/walker"
)

const testFile = "walker_test"

func TestWalker(t *testing.T) {
	t.Run("current_dir", func(tt *testing.T) {
		testWalker(tt, ".")
	})

	t.Run("parent_recursive", func(tt *testing.T) {
		testWalker(tt, "../...")
	})
}

func testWalker(t *testing.T, path string) {
	var fset = token.NewFileSet()
	var w = walker.NewWalker(fset)

	var err = w.Walk(path)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	w.Close()

	var gotNames = make(map[string]struct{})
	for f := range w.Files() {
		gotNames[f.Name.String()] = struct{}{}
	}

	if len(gotNames) == 0 {
		t.Fatal("got no names")
	}

	_, gotWalkerTest := gotNames[testFile]
	if !gotWalkerTest {
		t.Fatalf("%s: not found in list: %+v", testFile, gotNames)
	}
}
