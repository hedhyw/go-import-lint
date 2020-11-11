# go-import-lint

![Version](https://img.shields.io/github/v/tag/hedhyw/go-import-lint)
[![Build Status](https://travis-ci.org/hedhyw/go-import-lint.svg?branch=master)](https://travis-ci.org/hedhyw/go-import-lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/go-import-lint)](https://goreportcard.com/report/github.com/hedhyw/go-import-lint)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/go-import-lint/badge.svg?branch=master)](https://coveralls.io/github/hedhyw/go-import-lint?branch=master)

Golang source code analyzer that checks imports order. It verifies that:
- standard, current package, and vendor imports are separated by a line;
- there are no blank lines between one import group;
- there are no more than two lines.

Example of good imports order:

<!-- ReadmeExample -->
```go
package main

import (
    // Standart imports.
    "fmt"
    "error"

    // Current package imports.
    "github.com/hedhyw/go-import-lint/internal/linter"
    "github.com/hedhyw/go-import-lint/internal/model"

    // External imports.
    "github.com/hedhyw/jsonscjson"
    "github.com/stretchr/testify/assert"

    // Unused imports.
    _ "github.com/lib/pq"
)
```
<!-- /ReadmeExample -->

## Features:

- [x] Checking the blank lines between standart, package and external imports.
- [x] Files discovering.
- [x] Support import aliases.
- [x] Getting package name from the **go.mod**.
- [x] Ignore vendor by default.
- [x] Ignore generated files.
- [x] Support comments offset.
- [x] 😠 Nolint comment `// nolint:go-import-lint`.
- [ ] Check imports arrange.
- [ ] Integrate with `golangci-lint`.

## Installing:

```sh
go get github.com/hedhyw/go-import-lint/cmd/go-import-lint
```

## Usage example:

Run:

`go-import-lint`

```
Usage of go-import-lint:
  -exclude value
        paths to exclude (default ./vendor, ./.git)
  -path value
        paths to lint (default ./...)
  -pkg string
        module package
```