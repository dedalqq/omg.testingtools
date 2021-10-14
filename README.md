# omg.testingTools

[![Go](https://github.com/dedalqq/omg.testingtools/actions/workflows/go.yml/badge.svg)](https://github.com/dedalqq/omg.testingtools/actions/workflows/go.yml)

This tool can be useful for writing a tests. If you want change private field in struct from imported libraries than it can help you.

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
