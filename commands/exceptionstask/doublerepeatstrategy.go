package exceptionstask

import (
	"github.com/bindasov/spaceBattle/commands"
)

type DoubleRepeatStrategy interface {
	Execute()
	CacheCommand(command commands.StraightMoveCommand)
}

func NewDoubleRepeatStrategy(
	log LogCommand,
	doubleRepeatCommand DoubleRepeatCommand,
) DoubleRepeatStrategy {
	doubleRepeatStrategy := &doubleRepeatStrategy{
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
		s.doubleRepeatCommand.CacheCommand(s.command)
		if err := s.doubleRepeatCommand.Execute(); err != nil {
			s.log.SetError(err)
			s.log.Execute()
		}
	}
}

func (s *doubleRepeatStrategy) CacheCommand(command commands.StraightMoveCommand) {
	s.command = command
}
