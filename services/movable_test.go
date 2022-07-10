package services

import (
	"spaceBattle/models"
	"spaceBattle/services/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove_Execute(t *testing.T) {
	type deps struct {
		movableMock *mocks.Movable
		moveService *Move
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "success",
			handler: func(t *testing.T, deps *deps) {
				currentPosition := &models.Vector{
					X: 12,
					Y: 5,
				}
				velocity := &models.Vector{
					X: -7,
					Y: 3,
				}

				deps.movableMock.On("GetPosition").Return(currentPosition)
				deps.movableMock.On("GetVelocity").Return(velocity)

				newPosition := deps.moveService.SumVectors(currentPosition, velocity)
				deps.movableMock.On("SetPosition", newPosition).Return(newPosition)

				result := deps.moveService.Execute()
				desiredPosition := &models.Vector{
					X: 5,
					Y: 8,
				}
				require.Equal(t, desiredPosition, result)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			movableMock := mocks.NewMovable(t)
			story := NewMove(movableMock)

			deps := &deps{
				movableMock: movableMock,
				moveService: story,
			}

			tc.handler(t, deps)
		})
	}
}
