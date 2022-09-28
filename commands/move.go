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
	currentPosition, err := m.movableAdapter.GetPosition()
	if err != nil {
		return err
	}

	velocity, err := m.movableAdapter.GetVelocity()
	if err != nil {
		return err
	}

	newPosition := m.sumVectors(currentPosition, velocity)
	if err := m.movableAdapter.SetPosition(newPosition); err != nil {
		return err
	}
	return nil
}

func (m *moveCommand) sumVectors(vec1 *models.Vector, vec2 *models.Vector) *models.Vector {
	return &models.Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}
