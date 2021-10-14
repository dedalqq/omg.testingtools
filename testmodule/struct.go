package testmodule

type TestStruct struct {
	a bool
	b int
	c string
	d []int
}

func NewTestStruct(a bool, b int, c string, d []int) TestStruct {
	return TestStruct{
		a: a, b: b, c: c, d: d,
	}
}
