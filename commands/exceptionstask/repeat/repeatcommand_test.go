package repeat

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/commands/exceptionstask/base"
	commandMock "github.com/bindasov/spaceBattle/commands/exceptionstask/base/mocks"
)

func TestRepeatCommand_Execute(t *testing.T) {
	type deps struct {
		command       *commandMock.Command
		repeatCommand *RepeatCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				deps.command.On("Execute").Return(nil)
				err := deps.repeatCommand.Execute()

				require.NoError(t, err)
			},
		},
		{
			name: "error",
			handler: func(t *testing.T, deps *deps) {
				executeErr := errors.New("error")
				deps.command.On("Execute").Return(executeErr)
				err := deps.repeatCommand.Execute()

				require.Errorf(t, err, base.RepeatCommandError.Error())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			command := &commandMock.Command{}
			repeatCommand := NewRepeatCommand(command)

			deps := &deps{
				command:       command,
				repeatCommand: repeatCommand,
			}

			tc.handler(t, deps)
		})
	}
}
