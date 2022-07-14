package services

type Rotable interface {
	GetProperty(key string) (interface{}, error)
	SetProperty(key string, p interface{}) error
}
