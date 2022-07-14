package services

import (
	"errors"
)

var (
	PositionPropertyNotFound         = errors.New("position property not found")
	VelocityPropertyNotFound         = errors.New("velocity property not found")
	DirectionPropertyNotFound        = errors.New("direction property not found")
	DirectionsNumberPropertyNotFound = errors.New("directions number property not found")

	PositionPropertySetError = errors.New("position property set error")

	PositionPropertyConversionError         = errors.New("position property conversion error")
	VelocityPropertyConversionError         = errors.New("velocity property conversion error")
	DirectionPropertyConversionError        = errors.New("position property conversion error")
	DirectionsNumberPropertyConversionError = errors.New("velocity property conversion error")
)
