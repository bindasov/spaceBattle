package multithreading

import "github.com/bindasov/spaceBattle/commands"

type SoftStopCommand interface {
	Execute() error
}

func NewSoftStop(commands chan commands.Command) SoftStopCommand {
	softStopCommand := &softStopCommand{
		commands: commands,
	}
	return softStopCommand
}

type softStopCommand struct {
	commands chan commands.Command
}

func (c *softStopCommand) Execute() error {
	close(c.commands)
	return nil
}
