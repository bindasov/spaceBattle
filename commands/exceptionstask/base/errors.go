package base

import (
	"errors"
)

var (
	CommandError             = errors.New("error command")
	RepeatCommandError       = errors.New("error repeat command")
	DoubleRepeatCommandError = errors.New("error double repeat command")
)
