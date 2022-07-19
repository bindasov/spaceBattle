package commands

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/adapters/mocks"
	"github.com/bindasov/spaceBattle/models"
)

func TestSpaceship_Execute(t *testing.T) {
	type deps struct {
		movableAdapterMock *mocks.MovableAdapter
		moveCommand        MoveCommand
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "execute success",
			handler: func(t *testing.T, deps *deps) {
				var currentPosition = &models.Vector{
					X: 12,
					Y: 5,
				}

				var velocity = &models.Vector{
					X: -7,
					Y: 3,
				}

				desiredPosition := &models.Vector{
					X: 5,
					Y: 8,
				}

				deps.movableAdapterMock.On("GetPosition").Return(currentPosition, nil)
				deps.movableAdapterMock.On("GetVelocity").Return(velocity, nil)
				deps.movableAdapterMock.On("SetPosition", desiredPosition).Return(nil)
				result, err := deps.moveCommand.Execute()
				require.Equal(t, desiredPosition, result)
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			movableAdapterMock := mocks.NewMovableAdapter(t)
			moveCommand := NewMove(movableAdapterMock)

			deps := &deps{
				movableAdapterMock: movableAdapterMock,
				moveCommand:        moveCommand,
			}

			tc.handler(t, deps)
		})
	}
}
