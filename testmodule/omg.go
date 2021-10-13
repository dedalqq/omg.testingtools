package testmodule

type TestStruct struct {
	a int
	b string
	c []int
}

func NewQQ(a int, b string, c []int) TestStruct {
	return TestStruct{
		a, b, c,
	}
}
