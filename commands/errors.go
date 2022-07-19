package commands

import (
	"errors"
)

var (
	CheckFuelCommandException = errors.New("check fuel command exception")
	CommandException          = errors.New("command exception")
)
