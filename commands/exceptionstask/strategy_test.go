package exceptionstask

import (
	"errors"
	"testing"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/mocks"
	moveCommandMock "github.com/bindasov/spaceBattle/commands/mocks"
)

func TestStrategy_Execute(t *testing.T) {
	type deps struct {
		command       *moveCommandMock.StraightMoveCommand
		logCommand    *mocks.LogCommand
		repeatCommand *mocks.RepeatCommand
		strategy      Strategy
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "full",
			handler: func(t *testing.T, deps *deps) {
				err := errors.New("error")

				deps.command.On("Execute").Return(err)

				deps.repeatCommand.On("CacheCommand", deps.command)
				deps.repeatCommand.On("Execute").Return(err)

				deps.logCommand.On("SetError", err)
				deps.logCommand.On("Execute")

				deps.strategy.Execute()
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logCommand := &mocks.LogCommand{}
			repeatCommand := &mocks.RepeatCommand{}
			moveCommand := &moveCommandMock.StraightMoveCommand{}
			strategy := NewStrategy(moveCommand, logCommand, repeatCommand)

			deps := &deps{
				command:       moveCommand,
				logCommand:    logCommand,
				repeatCommand: repeatCommand,
				strategy:      strategy,
			}

			tc.handler(t, deps)
		})
	}
}
