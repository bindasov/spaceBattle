package base

import "github.com/bindasov/spaceBattle/commands"

func NewCommand() commands.Command {
	command := &command{}
	return command
}

type command struct {
}

func (c *command) Execute() error {
	return CommandError
}
