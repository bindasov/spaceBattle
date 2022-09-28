package commands

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/commands/mocks"
)

func TestStraightMoveCommand_Execute(t *testing.T) {
	type deps struct {
		checkFuelMock *mocks.CheckFuelCommand
		burnFuelMock  *mocks.BurnFuelCommand
		moveMock      *mocks.MoveCommand
		straightMove  StraightMoveCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				deps.checkFuelMock.On("Execute").Return(nil)
				deps.moveMock.On("Execute").Return(nil)
				deps.burnFuelMock.On("Execute").Return(nil)
				err := deps.straightMove.Execute()
				require.Nil(t, err)
			},
		},
		{
			name: "check fuel command exception",
			handler: func(t *testing.T, deps *deps) {
				deps.checkFuelMock.On("Execute").Return(CheckFuelCommandException)
				err := deps.straightMove.Execute()
				require.Errorf(t, err, CommandException.Error())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			checkFuelMock := mocks.NewCheckFuelCommand(t)
			burnFuelMock := mocks.NewBurnFuelCommand(t)
			moveMock := &mocks.MoveCommand{}
			straightMove := NewStraightMove(checkFuelMock, burnFuelMock, moveMock)

			deps := &deps{
				checkFuelMock: checkFuelMock,
				burnFuelMock:  burnFuelMock,
				moveMock:      moveMock,
				straightMove:  straightMove,
			}

			tc.handler(t, deps)
		})
	}
}
