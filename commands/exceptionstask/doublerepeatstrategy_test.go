package exceptionstask

import (
	"errors"
	"github.com/bindasov/spaceBattle/commands/exceptionstask/mocks"
	"testing"

	moveCommandMock "github.com/bindasov/spaceBattle/commands/mocks"
)

func TestDoubleRepeatStrategy_Execute(t *testing.T) {
	type deps struct {
		command              *moveCommandMock.StraightMoveCommand
		logCommand           *mocks.LogCommand
		doubleRepeatCommand  *mocks.DoubleRepeatCommand
		doubleRepeatStrategy DoubleRepeatStrategy
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "full",
			handler: func(t *testing.T, deps *deps) {
				command := &moveCommandMock.StraightMoveCommand{}
				err := errors.New("error")
				deps.doubleRepeatStrategy.CacheCommand(command)
				command.On("Execute").Return(err)

				deps.doubleRepeatCommand.On("CacheCommand", command)
				deps.doubleRepeatCommand.On("Execute").Return(err)

				deps.logCommand.On("SetError", err)
				deps.logCommand.On("Execute")

				deps.doubleRepeatStrategy.Execute()
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logCommand := &mocks.LogCommand{}
			doubleRepeatCommand := &mocks.DoubleRepeatCommand{}
			strategy := NewDoubleRepeatStrategy(logCommand, doubleRepeatCommand)

			deps := &deps{
				logCommand:           logCommand,
				doubleRepeatCommand:  doubleRepeatCommand,
				doubleRepeatStrategy: strategy,
			}

			tc.handler(t, deps)
		})
	}
}
