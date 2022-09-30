package spaceship

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckFuelCommand_Execute(t *testing.T) {
	type deps struct {
		command CheckFuelCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				err := deps.command.Execute()
				require.NoError(t, err)
			},
		},
		{
			name: "command error",
			handler: func(t *testing.T, deps *deps) {
				deps.command.SetLeftover(0)
				err := deps.command.Execute()
				require.Errorf(t, err, CheckFuelCommandException.Error())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commandMock := NewCheckFuel(5, 5)

			deps := &deps{
				command: commandMock,
			}

			tc.handler(t, deps)
		})
	}
}
