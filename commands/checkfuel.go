package commands

type CheckFuelCommand interface {
	Execute() error
	SetLeftover(leftover int)
}

func NewCheckFuel(leftover int, consumption int) CheckFuelCommand {
	checkFuelCommand := &checkFuelCommand{
		leftover:    leftover,
		consumption: consumption,
	}
	return checkFuelCommand
}

type checkFuelCommand struct {
	leftover    int
	consumption int
}

func (c *checkFuelCommand) Execute() error {
	if c.leftover < c.consumption {
		return CheckFuelCommandException
	}
	c.SetLeftover(c.leftover - c.consumption)
	return nil
}

func (c *checkFuelCommand) SetLeftover(leftover int) {
	c.leftover = leftover
}
