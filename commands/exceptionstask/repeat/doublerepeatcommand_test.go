package repeat

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
	commandMock "github.com/bindasov/spaceBattle/commands/exceptionstask/base/mocks"
)

func TestDoubleRepeatCommand_Execute(t *testing.T) {
	type deps struct {
		command             *commandMock.Command
		doubleRepeatCommand *DoubleRepeatCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				deps.command.On("Execute").Return(nil)
				err := deps.doubleRepeatCommand.Execute()

				require.NoError(t, err)
			},
		},
		{
			name: "error",
			handler: func(t *testing.T, deps *deps) {
				executeErr := errors.New("error")
				deps.command.On("Execute").Return(executeErr)
				err := deps.doubleRepeatCommand.Execute()

				require.Errorf(t, err, base.DoubleRepeatCommandError.Error())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			command := &commandMock.Command{}
			doubleRepeatCommand := NewDoubleRepeatCommand(command)

			deps := &deps{
				command:             command,
				doubleRepeatCommand: doubleRepeatCommand,
			}

			tc.handler(t, deps)
		})
	}
}
