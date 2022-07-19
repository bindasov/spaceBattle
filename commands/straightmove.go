package commands

type StraightMoveCommand interface {
	Execute() error
}

func NewStraightMove(
	checkFuelCommand CheckFuelCommand, burnFuelCommand BurnFuelCommand, moveCommand MoveCommand) StraightMoveCommand {
	straightMoveCommand := &straightMoveCommand{
		checkFuelCommand: checkFuelCommand,
		burnFuelCommand:  burnFuelCommand,
		moveCommand:      moveCommand,
	}
	return straightMoveCommand
}

type straightMoveCommand struct {
	checkFuelCommand CheckFuelCommand
	burnFuelCommand  BurnFuelCommand
	moveCommand      MoveCommand
}

func (c *straightMoveCommand) Execute() error {
	if err := c.checkFuelCommand.Execute(); err != nil {
		return CommandException
	}
	_, err := c.moveCommand.Execute()
	if err != nil {
		return CommandException
	}
	if err := c.burnFuelCommand.Execute(); err != nil {
		return CommandException
	}
	return nil
}
