package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestStringsFlag(t *testing.T) {
	expValues := []string{"1", "1"}

	f := stringsFlag{
		values:    []string{"default"},
		isInitial: true,
	}

	fs := flag.NewFlagSet("testset", flag.PanicOnError)
	fs.Var(&f, "test", "")

	err := fs.Parse([]string{"-test", "1", "-test", "1"})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(expValues, f.values) {
		t.Fatalf("values: exp %+v, got %+v", expValues, f.values)
	}
}

func TestMain(t *testing.T) {
	flags := newFlags()

	err := flags.Paths.Set("../../...")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	code := run(flags)
	if code != 0 {
		t.Fatalf("code: exp 0, got: %d", code)
	}
}
