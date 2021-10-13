package testingtools

import (
	"fmt"
	"omg.testingtools/testmodule"
	"testing"
)

func TestSetPrivateValue(t *testing.T) {
	qq := testmodule.NewQQ(123, "abc", []int{5, 6, 7})

	SetPrivateValue(&qq, "c", []int{1, 2, 3, 4})
	SetPrivateValue(&qq, "b", "qwe")

	if fmt.Sprint(qq) != "{123 qwe [1 2 3 4]}" {
		t.Fail()
	}
}
