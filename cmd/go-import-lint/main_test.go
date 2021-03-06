package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestStringsFlag(t *testing.T) {
	var expValues = []string{"1", "1"}

	var f = stringsFlag{
		values:    []string{"default"},
		isInitial: true,
	}

	var fs = flag.NewFlagSet("testset", flag.PanicOnError)
	fs.Var(&f, "test", "")

	var err = fs.Parse([]string{"-test", "1", "-test", "1"})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(expValues, f.values) {
		t.Fatalf("values: exp %+v, got %+v", expValues, f.values)
	}
}

func TestMain(t *testing.T) {
	var flags = newFlags()

	var err = flags.Paths.Set("../../...")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	var code = run(flags)
	if code != 0 {
		t.Fatalf("code: exp 0, got: %d", code)
	}
}
