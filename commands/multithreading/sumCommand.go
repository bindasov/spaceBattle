package multithreading

type TestCommand interface {
	Execute() error
	GetResult() int
}

func NewSum(a, b int) TestCommand {
	testCommand := &testCommand{
		a: a,
		b: b,
	}
	return testCommand
}

type testCommand struct {
	a   int
	b   int
	sum int
}

func (c *testCommand) Execute() error {
	c.sum = c.a + c.b
	return nil
}

func (c *testCommand) GetResult() int {
	return c.sum
}
