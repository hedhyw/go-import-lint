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

## TODO:
**Currently the work in progress.**

Current tasks:
- [x] Minimal application.
- [x] Check blank lines between imports.
- [x] Discover files.
- [x] Support comments.
- [ ] Look at integration with other linters.
- [ ] Check linter order.
- [ ] Get package from **go.mod** if possible.
- [ ] Ignore generated files.
- [ ] Ignore vendor.
- [ ] 😠 Nolint comment.
- [x] Run `go-import-lint` for this repository.
- [x] Parse readme program in test.
- [ ] Documentate.
- [ ] Increase test coverage.

## Installing:

```sh
go get github.com/hedhyw/go-import-lint
```

## Usage example:

Run:

`go-import-lint -path ./... -pkg PACKAGE_URL`
