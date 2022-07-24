package space_battle

import (
	"github.com/bindasov/spaceBattle/commands"
)

type spaceship struct {
	moveCommand commands.MoveCommand
}

func (m *spaceship) Execute() {
	if err := m.moveCommand.Execute(); err != nil {
		return
	}
}
