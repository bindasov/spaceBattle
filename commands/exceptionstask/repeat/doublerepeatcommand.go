package repeat

import (
	"github.com/bindasov/spaceBattle/commands"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
)

func NewDoubleRepeatCommand(command commands.Command) *DoubleRepeatCommand {
	doubleRepeatCommand := &DoubleRepeatCommand{
		command: command,
	}
	return doubleRepeatCommand
}

type DoubleRepeatCommand struct {
	command commands.Command
}

func (rc *DoubleRepeatCommand) Execute() error {
	if err := rc.command.Execute(); err != nil {
		if err := rc.command.Execute(); err != nil {
			if err := rc.command.Execute(); err != nil {
				return base.DoubleRepeatCommandError
			}
		}
	}
	return nil
}
