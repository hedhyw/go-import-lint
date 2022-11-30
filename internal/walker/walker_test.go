package walker_test

import (
	"go/token"
	"testing"

	"github.com/hedhyw/go-import-lint/internal/walker"
)

const thisFile = "walker_test"

func TestWalker(t *testing.T) {
	t.Run("current_dir", func(tt *testing.T) {
		testWalkerThisFileFound(tt, ".")
	})

	t.Run("current_file", func(tt *testing.T) {
		testWalkerThisFileFound(tt, thisFile+".go")
	})

	t.Run("parent_recursive", func(tt *testing.T) {
		testWalkerThisFileFound(tt, "../...")
	})
}

func testWalkerThisFileFound(t *testing.T, path string) {
	gotNames := getWalkerResult(t, path, nil)

	if len(gotNames) == 0 {
		t.Fatal("got no names")
	}

	_, gotWalkerTest := gotNames[thisFile]
	if !gotWalkerTest {
		t.Fatalf("%s: not found in list: %+v", thisFile, gotNames)
	}
}

func TestWalkerExclude(t *testing.T) {
	t.Run("directory", func(tt *testing.T) {
		testExcludeThisFile(tt, []string{"../walker"})
	})

	t.Run("file", func(tt *testing.T) {
		testExcludeThisFile(tt, []string{thisFile + ".go"})
	})
}

func testExcludeThisFile(t *testing.T, exclude []string) {
	gotNames := getWalkerResult(t, "../...", exclude)

	if len(gotNames) == 0 {
		t.Fatal("got no names")
	}

	_, gotWalkerTest := gotNames[thisFile]
	if gotWalkerTest {
		t.Fatalf("%s: found in list: %+v", thisFile, gotNames)
	}
}

func TestWalkerIgnoreHandleNoGoFile(t *testing.T) {
	_ = getWalkerResult(t, "../../go.mod", nil)
}

func TestWalkerErrors(t *testing.T) {
	t.Run("file-not-found", func(tt *testing.T) {
		testWalkerError(tt, "not-found.go")
	})

	t.Run("dir-not-found", func(tt *testing.T) {
		testWalkerError(tt, ".../...")
	})
}

func testWalkerError(t *testing.T, path string) {
	fset := token.NewFileSet()

	w, err := walker.NewWalker(fset, []string{})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if err = w.Walk(path); err == nil {
		t.Fatalf("err: exp not exist, got %s", err)
	}
}

func getWalkerResult(t *testing.T, path string, exclude []string) (gotNames map[string]struct{}) {
	fset := token.NewFileSet()

	w, err := walker.NewWalker(fset, exclude)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if err = w.Walk(path); err != nil {
		t.Fatalf("err: %s", err)
	}

	w.Close()

	gotNames = make(map[string]struct{})
	for f := range w.Files() {
		gotNames[f.Name.String()] = struct{}{}
	}

	return gotNames
}
