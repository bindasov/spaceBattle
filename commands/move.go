package commands

import (
	"github.com/bindasov/spaceBattle/adapters"
	"github.com/bindasov/spaceBattle/models"
)

type MoveCommand interface {
	Execute() error
}

func NewMove(movableAdapter adapters.MovableAdapter) MoveCommand {
	moveCommand := &moveCommand{movableAdapter: movableAdapter}
	return moveCommand
}

type moveCommand struct {
	movableAdapter adapters.MovableAdapter
}

func (m *moveCommand) Execute() error {
	currentPosition := m.movableAdapter.GetPosition()

	velocity := m.movableAdapter.GetVelocity()

	newPosition := m.sumVectors(currentPosition, velocity)
	m.movableAdapter.SetPosition(newPosition)
	return nil
}

func (m *moveCommand) sumVectors(vec1 *models.Vector, vec2 *models.Vector) *models.Vector {
	return &models.Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}
