# omg.testingTools

[![Go](https://github.com/dedalqq/omg.testingtools/actions/workflows/go.yml/badge.svg)](https://github.com/dedalqq/omg.testingtools/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dedalqq/omg.testingtools)](https://goreportcard.com/report/github.com/dedalqq/omg.testingtools)
[![Coverage Status](https://coveralls.io/repos/github/dedalqq/omg.testingtools/badge.svg?branch=master)](https://coveralls.io/github/dedalqq/omg.testingtools?branch=master)

This tool can be useful for writing a tests. If you want change private field in struct from imported libraries than it can help you.

## Example

```go
package main

import (
	"strings"
	
	"github.com/dedalqq/omg.testingtools"
)

func TestSomeCase(t *testing.T) {
	buffer := strings.Builder{}
	
	testingtools.SetPrivateValue(&buffer, "buf", []byte("data in strings buffer"))
}
```
## How it works

This function finding needs field via `reflect` then get pointer from field address via `unsafe.Pointer` and write value via variable with the same address as original struct field.
