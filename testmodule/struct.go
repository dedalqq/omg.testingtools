package testmodule

// TestStruct an example of structure for test
type TestStruct struct {
	a bool
	b int
	c string
	d []int
}

// NewTestStruct returns new TestStruct for tests
func NewTestStruct(a bool, b int, c string, d []int) TestStruct {
	return TestStruct{
		a: a, b: b, c: c, d: d,
	}
}
