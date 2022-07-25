package exceptionstask

import (
	"github.com/bindasov/spaceBattle/commands"
)

type RepeatCommand interface {
	Execute() error
	CacheCommand(command commands.StraightMoveCommand)
}

func NewRepeatCommand() RepeatCommand {
	repeatCommand := &repeatCommand{}
	return repeatCommand
}

type repeatCommand struct {
	command commands.StraightMoveCommand
}

func (rc *repeatCommand) Execute() error {
	if err := rc.command.Execute(); err != nil {
		return err
	}
	return nil
}
func (rc *repeatCommand) CacheCommand(command commands.StraightMoveCommand) {
	rc.command = command
}
