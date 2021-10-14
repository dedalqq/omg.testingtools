# omg.testingTools

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