package base

type Command interface {
	Execute() error
}

func NewCommand() Command {
	command := &command{}
	return command
}

type command struct {
}

func (c *command) Execute() error {
	return CommandError
}
