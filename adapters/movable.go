package adapters

import (
	"math"

	"github.com/bindasov/spaceBattle/models"
	"github.com/bindasov/spaceBattle/services"
)

type MovableAdapter interface {
	GetPosition() (*models.Vector, error)
	GetVelocity() (*models.Vector, error)
	SetPosition(position *models.Vector) error
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

func (m *movableAdapter) GetPosition() (*models.Vector, error) {
	value, err := m.movable.GetProperty("Position")
	if err != nil {
		return nil, services.PositionPropertyNotFound
	}
	position, ok := value.(*models.Vector)
	if ok == false {
		return nil, services.PositionPropertyConversionError
	}
	return position, nil
}

func (m *movableAdapter) GetVelocity() (*models.Vector, error) {
	direction, err := m.getDirection()
	if err != nil {
		return nil, err
	}

	directionsNumber, err := m.getDirectionsNumber()
	if err != nil {
		return nil, err
	}

	velocityValue, err := m.getVelocity()
	if err != nil {
		return nil, err
	}

	return &models.Vector{
		X: int(float64(velocityValue) * math.Cos(float64(direction)/float64(360*directionsNumber))),
		Y: int(float64(velocityValue) * math.Sin(float64(direction)/float64(360*directionsNumber))),
	}, nil
}

func (m *movableAdapter) SetPosition(position *models.Vector) error {
	var pos interface{} = position
	if err := m.movable.SetProperty("Position", pos); err != nil {
		return services.PositionPropertySetError
	}
	return nil
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
