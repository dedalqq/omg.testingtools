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

func TestNotPointerError(t *testing.T) {
	defer func() {
		if r := recover(); r != "object is not a pointer" {
			t.Fail()
		}
	}()

	qq := testmodule.NewTestStruct(true, 123, "abc", []int{5, 6, 7})

	SetPrivateValue(qq, "a", false)
}

func TestNotStructError(t *testing.T) {
	defer func() {
		if r := recover(); r != "object is not a pointer to struct" {
			fmt.Println(r)
			t.Fail()
		}
	}()

	notObject := "string"

	SetPrivateValue(&notObject, "a", false)
}

func TestIncorrectValueTypeError(t *testing.T) {
	defer func() {
		if r := recover(); r != "incorrect type: must be []int" {
			t.Fail()
		}
	}()

	qq := testmodule.NewTestStruct(true, 123, "abc", []int{5, 6, 7})

	SetPrivateValue(&qq, "d", []byte{5, 6, 7})
}

func TestIncorrectFieldNameError(t *testing.T) {
	defer func() {
		if r := recover(); r != "Field [incorrectValue] not found in struct" {
			t.Fail()
		}
	}()

	qq := testmodule.NewTestStruct(true, 123, "abc", []int{5, 6, 7})

	SetPrivateValue(&qq, "incorrectValue", false)
}
