package multithreading

import (
	"fmt"
	"github.com/bindasov/spaceBattle/commands"
)

type StartCommand interface {
	Execute() error
}

func NewStart(commands chan commands.Command) StartCommand {
	startCommand := &startCommand{
		commands: commands,
	}
	return startCommand
}

type startCommand struct {
	commands chan commands.Command
}

func (c *startCommand) Execute() error {
	for command := range c.commands {
		fmt.Println("executing actor command")
		if err := command.Execute(); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
