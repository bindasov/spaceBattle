package space_battle

import (
	"github.com/bindasov/spaceBattle/adapters"
	"github.com/bindasov/spaceBattle/models"
)

type spaceship struct {
	movableAdapter adapters.MovableAdapter
}

func (m *spaceship) SumVectors(vec1 *models.Vector, vec2 *models.Vector) *models.Vector {
	return &models.Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}

func (m *spaceship) Execute() (*models.Vector, error) {
	currentPosition, err := m.movableAdapter.GetPosition()
	if err != nil {
		return nil, err
	}

	velocity, err := m.movableAdapter.GetVelocity()
	if err != nil {
		return nil, err
	}

	newPosition := m.SumVectors(currentPosition, velocity)
	if err := m.movableAdapter.SetPosition(newPosition); err != nil {
		return nil, err
	}
	return newPosition, nil
}
