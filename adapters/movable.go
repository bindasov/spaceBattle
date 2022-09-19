package adapters

import (
	"math"

	"github.com/bindasov/spaceBattle/models"
	"github.com/bindasov/spaceBattle/services"
)

type MovableAdapter interface {
	GetPosition() *models.Vector
	GetVelocity() *models.Vector
	SetPosition(position *models.Vector)
}

func NewMovable(movable services.Movable, rotable services.Rotable) MovableAdapter {
	movableAdapter := &movableAdapter{
		movable: movable,
		rotable: rotable,
	}
	return movableAdapter
}

type movableAdapter struct {
	movable services.Movable
	rotable services.Rotable
}

func (m *movableAdapter) GetPosition() *models.Vector {
	value, _ := m.movable.GetProperty("Position")
	position, _ := value.(*models.Vector)
	return position
}

func (m *movableAdapter) GetVelocity() *models.Vector {
	direction, _ := m.getDirection()

	directionsNumber, _ := m.getDirectionsNumber()

	velocityValue, _ := m.getVelocity()

	return &models.Vector{
		X: int(float64(velocityValue) * math.Cos(float64(direction)/float64(360*directionsNumber))),
		Y: int(float64(velocityValue) * math.Sin(float64(direction)/float64(360*directionsNumber))),
	}
}

func (m *movableAdapter) SetPosition(position *models.Vector) {
	var pos interface{} = position
	_ = m.movable.SetProperty("Position", pos)
}

func (m *movableAdapter) getDirection() (int, error) {
	value, err := m.rotable.GetProperty("Direction")
	if err != nil {
		return 0, services.DirectionPropertyNotFound
	}

	direction, ok := value.(int)
	if ok == false {
		return 0, services.DirectionPropertyConversionError
	}
	return direction, nil
}

func (m *movableAdapter) getDirectionsNumber() (int, error) {
	value, err := m.rotable.GetProperty("DirectionsNumber")
	if err != nil {
		return 0, services.DirectionsNumberPropertyNotFound
	}

	directionsNumber, ok := value.(int)
	if ok == false {
		return 0, services.DirectionsNumberPropertyConversionError
	}
	return directionsNumber, nil
}

func (m *movableAdapter) getVelocity() (int, error) {
	value, err := m.movable.GetProperty("Velocity")
	if err != nil {
		return 0, services.VelocityPropertyNotFound
	}

	velocity, ok := value.(int)
	if ok == false {
		return 0, services.VelocityPropertyConversionError
	}
	return velocity, nil
}
