package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bindasov/spaceBattle/models"
	"github.com/bindasov/spaceBattle/services"
	"github.com/bindasov/spaceBattle/services/mocks"
)

func TestMovableAdapter_GetPosition(t *testing.T) {
	type deps struct {
		movableMock    *mocks.Movable
		rotableMock    *mocks.Rotable
		movableAdapter MovableAdapter
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "get position property not found error",
			handler: func(t *testing.T, deps *deps) {
				deps.movableMock.On("GetProperty", "Position").Return(nil,
					services.PositionPropertyNotFound)
				position := deps.movableAdapter.GetPosition()
				require.Nil(t, position)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			movableMock := mocks.NewMovable(t)
			rotableMock := mocks.NewRotable(t)
			adapter := NewMovable(movableMock, rotableMock)

			deps := &deps{
				movableMock:    movableMock,
				rotableMock:    rotableMock,
				movableAdapter: adapter,
			}

			tc.handler(t, deps)
		})
	}
}

func TestMovableAdapter_GetVelocity(t *testing.T) {
	type deps struct {
		movableMock    *mocks.Movable
		rotableMock    *mocks.Rotable
		movableAdapter MovableAdapter
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "get instant speed property not found error",
			handler: func(t *testing.T, deps *deps) {
				direction := 0
				directionsNumber := 0
				deps.rotableMock.On("GetProperty", "Direction").Return(direction, nil)
				deps.rotableMock.On("GetProperty", "DirectionsNumber").Return(directionsNumber, nil)
				deps.movableMock.On("GetProperty", "Velocity").Return(nil,
					services.VelocityPropertyNotFound)
				velocity := deps.movableAdapter.GetVelocity()
				require.Nil(t, velocity)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			movableMock := mocks.NewMovable(t)
			rotableMock := mocks.NewRotable(t)
			adapter := NewMovable(movableMock, rotableMock)

			deps := &deps{
				movableMock:    movableMock,
				rotableMock:    rotableMock,
				movableAdapter: adapter,
			}

			tc.handler(t, deps)
		})
	}
}

func TestMovableAdapter_SetPosition(t *testing.T) {
	type deps struct {
		movableMock    *mocks.Movable
		rotableMock    *mocks.Rotable
		movableAdapter MovableAdapter
	}
	tests := []struct {
		name    string
		handler func(*testing.T, *deps)
	}{
		{
			name: "set position property not found error",
			handler: func(t *testing.T, deps *deps) {
				var position interface{} = &models.Vector{}
				deps.movableMock.On("SetProperty", "Position", position).
					Return(services.PositionPropertySetError)
				deps.movableAdapter.SetPosition(&models.Vector{})
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			movableMock := mocks.NewMovable(t)
			rotableMock := mocks.NewRotable(t)
			adapter := NewMovable(movableMock, rotableMock)

			deps := &deps{
				movableMock:    movableMock,
				rotableMock:    rotableMock,
				movableAdapter: adapter,
			}

			tc.handler(t, deps)
		})
	}
}
