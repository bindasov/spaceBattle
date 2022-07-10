package services

type Rotable interface {
	GetDirection() int
	GetAngularVelocity() int
	GetDirectionsNumber() int
	SetDirection(d int) int
}

type Rotate struct {
	rotatableService Rotable
}

func NewRotate(
	iRotable Rotable,
) *Rotate {
	rotate := &Rotate{
		rotatableService: iRotable,
	}
	return rotate
}

func (r *Rotate) Execute() int {
	angularVelocity := r.rotatableService.GetAngularVelocity()
	directionsNumber := r.rotatableService.GetDirectionsNumber()
	angularRotation := angularVelocity % directionsNumber
	currentDirection := r.rotatableService.GetDirection()
	newDirection := r.rotatableService.SetDirection(currentDirection + angularRotation)
	return newDirection
}
