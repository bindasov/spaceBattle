package exceptionstask

import (
	"testing"

	"github.com/stretchr/testify/require"

	moveCommandMock "github.com/bindasov/spaceBattle/commands/mocks"
)

func TestDoubleRepeatCommand_Execute(t *testing.T) {
	type deps struct {
		doubleRepeatCommand DoubleRepeatCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				command := &moveCommandMock.StraightMoveCommand{}
				deps.doubleRepeatCommand.CacheCommand(command)

				command.On("Execute").Return(nil)

				err := deps.doubleRepeatCommand.Execute()

				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			command := NewDoubleRepeatCommand()

			deps := &deps{
				doubleRepeatCommand: command,
			}

			tc.handler(t, deps)
		})
	}
}
