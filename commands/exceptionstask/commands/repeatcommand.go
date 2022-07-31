package commands

import "github.com/bindasov/spaceBattle/commands/exceptionstask/base"

func NewRepeatCommand(command base.Command) *RepeatCommand {
	repeatCommand := &RepeatCommand{
		command: command,
	}
	return repeatCommand
}

type RepeatCommand struct {
	command base.Command
}

func (rc *RepeatCommand) Execute() error {
	if err := rc.command.Execute(); err != nil {
		return base.RepeatCommandError
	}
	return nil
}
