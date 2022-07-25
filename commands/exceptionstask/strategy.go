package exceptionstask

import (
	"github.com/bindasov/spaceBattle/commands"
)

type Strategy interface {
	Execute()
}

func NewStrategy(
	command commands.StraightMoveCommand,
	log LogCommand,
	repeatCommand RepeatCommand,
) Strategy {
	strategy := &strategy{
		command:       command,
		log:           log,
		repeatCommand: repeatCommand,
	}
	return strategy
}

type strategy struct {
	command       commands.StraightMoveCommand
	log           LogCommand
	repeatCommand RepeatCommand
}

func (s *strategy) Execute() {
	if err := s.command.Execute(); err != nil {
		s.repeatCommand.CacheCommand(s.command)
		if err := s.repeatCommand.Execute(); err != nil {
			s.log.SetError(err)
			s.log.Execute()
		}
	}
}
