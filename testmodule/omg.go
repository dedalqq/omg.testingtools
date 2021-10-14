package testmodule

type TestStruct struct {
	a bool
	b int
	c string
	d []int
}

func NewQQ(a int, b string, c []int) TestStruct {
	return TestStruct{
		false, a, b, c,
	}
}
