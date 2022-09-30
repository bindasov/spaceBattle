package repeat

import (
	"github.com/bindasov/spaceBattle/commands"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
)

func NewRepeatCommand(command commands.Command) *RepeatCommand {
	repeatCommand := &RepeatCommand{
		command: command,
	}
	return repeatCommand
}

type RepeatCommand struct {
	command commands.Command
}

func (rc *RepeatCommand) Execute() error {
	if err := rc.command.Execute(); err != nil {
		return base.RepeatCommandError
	}
	return nil
}
