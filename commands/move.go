package commands

import (
	"github.com/bindasov/spaceBattle/adapters"
	"github.com/bindasov/spaceBattle/models"
)

type MoveCommand interface {
	Execute() (*models.Vector, error)
}

func NewMove(movableAdapter adapters.MovableAdapter) MoveCommand {
	moveCommand := &moveCommand{movableAdapter: movableAdapter}
	return moveCommand
}

type moveCommand struct {
	movableAdapter adapters.MovableAdapter
}

func (m *moveCommand) Execute() (*models.Vector, error) {
	currentPosition, err := m.movableAdapter.GetPosition()
	if err != nil {
		return nil, err
	}

	velocity, err := m.movableAdapter.GetVelocity()
	if err != nil {
		return nil, err
	}

	newPosition := m.sumVectors(currentPosition, velocity)
	if err := m.movableAdapter.SetPosition(newPosition); err != nil {
		return nil, err
	}
	return newPosition, nil
}

func (m *moveCommand) sumVectors(vec1 *models.Vector, vec2 *models.Vector) *models.Vector {
	return &models.Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}
