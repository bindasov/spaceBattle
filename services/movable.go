package services

type Movable interface {
	GetProperty(key string) (interface{}, error)
	SetProperty(key string, p interface{}) error
}
