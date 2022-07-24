package exceptionstask

import (
	"github.com/bindasov/spaceBattle/commands"
)

type DoubleRepeatStrategy interface {
	Execute()
}

func NewDoubleRepeatStrategy(
	command commands.StraightMoveCommand,
	log LogCommand,
	doubleRepeatCommand DoubleRepeatCommand,
) DoubleRepeatStrategy {
	doubleRepeatStrategy := &doubleRepeatStrategy{
		command:             command,
		log:                 log,
		doubleRepeatCommand: doubleRepeatCommand,
	}
	return doubleRepeatStrategy
}

type doubleRepeatStrategy struct {
	command             commands.StraightMoveCommand
	log                 LogCommand
	doubleRepeatCommand DoubleRepeatCommand
}

func (s *doubleRepeatStrategy) Execute() {
	if err := s.command.Execute(); err != nil {
		s.doubleRepeatCommand.Set(s.command)
		if err := s.doubleRepeatCommand.Execute(); err != nil {
			s.log.SetError(err)
			s.log.Execute()
		}
	}
}
