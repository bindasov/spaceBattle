package services

import "spaceBattle/models"

type Movable interface {
	GetPosition() *models.Vector
	GetVelocity() *models.Vector
	SetPosition(v *models.Vector) *models.Vector
}

type Move struct {
	movableService Movable
}

func NewMove(
	iMovable Movable,
) *Move {
	move := &Move{
		movableService: iMovable,
	}
	return move
}

func (m *Move) SumVectors(vec1 *models.Vector, vec2 *models.Vector) *models.Vector {
	return &models.Vector{
		X: vec1.X + vec2.X,
		Y: vec1.Y + vec2.Y,
	}
}

func (m *Move) Execute() *models.Vector {
	currentPosition := m.movableService.GetPosition()
	velocity := m.movableService.GetVelocity()
	newPosition := m.SumVectors(currentPosition, velocity)
	return m.movableService.SetPosition(newPosition)
}
