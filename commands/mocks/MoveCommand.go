// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MoveCommand is an autogenerated mock type for the MoveCommand type
type MoveCommand struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *MoveCommand) Execute() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
