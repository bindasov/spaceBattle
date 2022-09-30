package spaceship

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBurnFuelCommand_Execute(t *testing.T) {
	type deps struct {
		command BurnFuelCommand
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commandMock := NewBurnFuel(5, 5)

			deps := &deps{
				command: commandMock,
			}

			tc.handler(t, deps)
		})
	}
}
