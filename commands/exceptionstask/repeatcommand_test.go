package exceptionstask

import (
	"testing"

	"github.com/stretchr/testify/require"

	moveCommandMock "github.com/bindasov/spaceBattle/commands/mocks"
)

func TestRepeatCommand_Execute(t *testing.T) {
	type deps struct {
		repeatCommand RepeatCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				command := &moveCommandMock.StraightMoveCommand{}
				deps.repeatCommand.CacheCommand(command)
				command.On("Execute").Return(nil)
				err := deps.repeatCommand.Execute()

				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repeatCommand := NewRepeatCommand()

			deps := &deps{
				repeatCommand: repeatCommand,
			}

			tc.handler(t, deps)
		})
	}
}
