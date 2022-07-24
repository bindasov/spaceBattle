package exceptionstask

import (
	"github.com/bindasov/spaceBattle/commands"
)

type DoubleRepeatCommand interface {
	Execute() error
	Set(command commands.StraightMoveCommand)
}

func NewDoubleRepeatCommand() DoubleRepeatCommand {
	doubleRepeatCommand := &doubleRepeatCommand{}
	return doubleRepeatCommand
}

type doubleRepeatCommand struct {
	command commands.StraightMoveCommand
}

func (drc *doubleRepeatCommand) Execute() error {
	if err := drc.command.Execute(); err != nil {
		if err = drc.command.Execute(); err != nil {
			return err
		}
	}
	return nil
}
func (drc *doubleRepeatCommand) Set(command commands.StraightMoveCommand) {
	drc.command = command
}
