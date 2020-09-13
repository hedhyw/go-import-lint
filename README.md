# go-import-lint

![Version](https://img.shields.io/github/v/tag/hedhyw/go-import-lint)
[![Build Status](https://travis-ci.org/hedhyw/go-import-lint.svg?branch=master)](https://travis-ci.org/hedhyw/go-import-lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/go-import-lint)](https://goreportcard.com/report/github.com/hedhyw/go-import-lint)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/go-import-lint/badge.svg?branch=master)](https://coveralls.io/github/hedhyw/go-import-lint?branch=master)

The linter checks that imports have a correct order.

Example of good order:

<!-- ReadmeExample -->
```go
package main

import (
    // Standart imports.
    "fmt"
    "error"

    // Current imports.
    "github.com/hedhyw/go-import-lint/internal"
    "github.com/hedhyw/go-import-lint/internal/model"

    // External imports.
    "github.com/hedhyw/jsonscjson"
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
go get github.com/hedhyw/go-import-lint
```

## Usage example:

Run:

`go-import-lint -path ./... -pkg PACKAGE_URL`

```
Usage of go-import-lint:
  -exclude value
        paths to exclude (default ./vendor, ./.git)
  -path value
        paths to lint (default ./...)
  -pkg string
        module package
```