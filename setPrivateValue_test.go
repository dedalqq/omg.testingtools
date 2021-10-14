package testingtools

import (
	"fmt"
	"testing"

	"omg.testingtools/testmodule"
)

func TestSetPrivateValue(t *testing.T) {
	qq := testmodule.NewTestStruct(true, 123, "abc", []int{5, 6, 7})

	SetPrivateValue(&qq, "a", false)
	SetPrivateValue(&qq, "d", []int{1, 2, 3, 4})
	SetPrivateValue(&qq, "c", "qwe")

	if fmt.Sprint(qq) != "{false 123 qwe [1 2 3 4]}" {
		t.Fail()
	}
}
