package space_battle

import (
	"github.com/bindasov/spaceBattle/commands"
	"github.com/bindasov/spaceBattle/models"
)

type spaceship struct {
	moveCommand commands.MoveCommand
}

func (m *spaceship) Execute() (*models.Vector, error) {
	return m.moveCommand.Execute()
}
